package middleware

import (
	"fmt"
	"go-gin-blog/internal/initializers"
	"go-gin-blog/internal/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
	//Get the cookie off req
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.Redirect(http.StatusTemporaryRedirect, "/login")
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	//Decode/validate it
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid{

		//check the exp
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.Redirect(http.StatusTemporaryRedirect, "/login")
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		//find the user with token sub
		var user models.User
		initializers.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			c.Redirect(http.StatusTemporaryRedirect, "/login")
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		 // Extract user ID from token claims
		 claims := token.Claims.(jwt.MapClaims)
		 userID := uint(claims["sub"].(float64))
 
		 // Set user ID in context
		 c.Set("user_id", userID)
 

		//attach to req
		c.Set("user", user)

		//continue
		c.Next()
	} else {
		c.Redirect(http.StatusTemporaryRedirect, "/login")
		c.AbortWithStatus(http.StatusUnauthorized)
	}

}
