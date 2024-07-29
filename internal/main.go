package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// r.LoadHTMLGlob("templates/*")
	// r.GET("/index", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "index.tmpl", gin.H{
	// 		"title": "Main website",
	// 	})
	// })
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
	r.Run()
}
