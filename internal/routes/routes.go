package routes

import (
	"go-gin-blog/internal/controllers"
	"go-gin-blog/internal/middleware"

	"github.com/gin-gonic/gin"
	// "github.com/jinzhu/gorm"
)

func SetupRoutes(r *gin.Engine) {


	r.GET("/signup", controllers.ShowSignupPage)
    r.GET("/login", controllers.ShowLoginPage)
	r.POST("/signupapi", controllers.Singup)
	r.POST("/loginapi", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.GET("/logout", controllers.Logout)

	authorized := r.Group("/")
    authorized.Use(middleware.RequireAuth)
    {
        authorized.GET("/home", controllers.Home)
    }
	
	// r.GET("/user", uc.GetUser)
	// r.GET("/logout", uc.GetUser)
    // r.POST("/login", uc.LoginUser)
	// r.PUT("/user/:id", uc.UpdateUser)

//     pc := &controllers.PostController{DB: db}

//     r.GET("/posts", pc.GetPosts)
//     r.POST("/posts", pc.CreatePost)
// 	r.PUT("/posts/:id", pc.UpdatePost)
// 	r.DELETE("/posts/:id", pc.DeletePost)

// 	cc := &controllers.CommentController{DB: db}

// 	r.GET("/comments", cc.GetComments)
//     r.POST("/comments", cc.CreateComment)
// 	r.PUT("/comments/:id", cc.UpdateComment)
// 	r.DELETE("/comments/:id", cc.DeleteComment)


}