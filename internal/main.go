package main

import (
	"fmt"
	"go-gin-blog/internal/initializers"
	"go-gin-blog/internal/routes"
	"net/http"

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
	// r.LoadHTMLGlob("templates/*")
	// r.GET("/index", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "index.tmpl", gin.H{
	// 		"title": "Main website",
	// 	})
	// })
	fmt.Println("hello world 2")
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the Home Page!",
		})
	})
	r.GET("/about", func(c *gin.Context) {
	    c.JSON(http.StatusOK, gin.H{
	        "message": "This is the About Page.",
	    })
	})
	r.GET("/user/:name", func(c *gin.Context) {
	    name := c.Param("name")
	    c.JSON(http.StatusOK, gin.H{
	        "message": "Hello " + name,
	    })
	})

	routes.SetupRoutes(r)
	r.Run()
}
