package main

import (
	"go-gin-blog/internal/controllers"
	"go-gin-blog/internal/initializers"
	"go-gin-blog/internal/middleware"
	// "go-gin-blog/internal/routes"

	"github.com/gin-gonic/gin"
)

func initialize() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	initialize()
	r := gin.Default()
	r.Static("./internal/assets", "./internal/assets")
	r.LoadHTMLGlob("./internal/templates/*")

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
		// Post routes
		authorized.GET("/posts", controllers.ShowPostsPage)
		// authorized.GET("/posts/:id", controllers.ShowViewPage)
		authorized.GET("/posts/create", controllers.ShowCreatePostPage)
		authorized.POST("/postsapi", controllers.CreatePost)
		authorized.GET("/posts/:id", controllers.GetPost)
		r.GET("/posts/:id/edit", controllers.EditPost)
		r.POST("/posts/:id", controllers.UpdatePost)
		authorized.DELETE("/posts/:id", controllers.DeletePost)
	}
	// routes.SetupRoutes(r)
	r.Run()
}
