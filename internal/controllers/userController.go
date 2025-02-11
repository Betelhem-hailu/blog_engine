package controllers

import (
	"go-gin-blog/internal/initializers"
	"net/http"
	"os"
	"time"

	"go-gin-blog/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// ShowSignupPage displays the signup page
func ShowSignupPage(c *gin.Context) {
    c.HTML(http.StatusOK, "signup.html", gin.H{
        "Title": "Signup",
    })
}

func Singup(c *gin.Context) {
	//Get the fullname email/pass req body
	var body struct {
        FullName string `form:"fullname"`
        Email    string `form:"email"`
        Password string `form:"password"`
    }

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}
	//Hash the paswrd
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})

		return
	}
	//create the user
	user := models.User{FullName: body.FullName, Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})

		return
	}

	//Respond
	c.Redirect(http.StatusMovedPermanently, "/login")
	// c.JSON(http.StatusOK, gin.H{})
}

// ShowLoginPage displays the login page
func ShowLoginPage(c *gin.Context) {
    c.HTML(http.StatusOK, "login.html", gin.H{
        "Title": "Login",
    })
}

func Login(c *gin.Context) {
	//Get the email and pass
	var body struct {
        Email    string `form:"email"`
        Password string `form:"password"`
    }

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	//Look up requested user
	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password invalid email",
		})

		return
	}

	//compare sent in pass with saved user pass hash
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password invalid password",
		})

		return
	}

	//Generate a jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid to create token",
		})

		return
	}

	//send it back
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	c.Redirect(http.StatusMovedPermanently, "/home")
	// c.JSON(http.StatusOK, gin.H{})
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}

func Logout(c *gin.Context) {
	c.SetCookie("Authorization", "", -1, "/", "localhost", false, true)
	c.Redirect(http.StatusSeeOther, "/login")
}

// Home displays the home page for logged-in users
func Home(c *gin.Context) {
	userID, exists := c.Get("user")
    if !exists {
        c.Redirect(http.StatusTemporaryRedirect, "/login")
        return
    }

    var user models.User
    initializers.DB.First(&user, userID)

    if user.ID == 0 {
        c.Redirect(http.StatusTemporaryRedirect, "/login")
        return
    }

    c.HTML(http.StatusOK, "home.html", gin.H{
        "Title": "Home",
        "User":  user,
    })
}
