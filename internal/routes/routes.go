package routes

import (
    "go-gin-blog/controllers"
    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
    pc := &controllers.PostController{DB: db}

    r.GET("/posts", pc.GetPosts)
    r.POST("/posts", pc.CreatePost)
	r.PUT("/posts/:id", pc.UpdatePost)
	r.DELETE("/posts/:id", pc.DeletePost)

	cc := &controllers.CommentController{DB: db}

	r.GET("/comments", cc.GetComments)
    r.POST("/comments", cc.CreateComment)
	r.PUT("/comments/:id", cc.UpdateComment)
	r.DELETE("/comments/:id", cc.DeleteComment)

	uc := &controllers.UserController{DB: db}

	r.GET("/user", uc.GetUser)
	r.GET("/logout", uc.GetUser)
    r.POST("/register", uc.CreateUser)
    r.POST("/login", uc.LoginUser)
	r.PUT("/user/:id", uc.UpdateUser)
}