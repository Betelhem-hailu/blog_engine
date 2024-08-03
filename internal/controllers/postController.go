package controllers

import (
	"fmt"
	"go-gin-blog/internal/initializers"
	"go-gin-blog/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreatePost creates a new post
func CreatePost(c *gin.Context) {
	// Parse form data
	var body struct {
		Title   string `form:"title" binding:"required"`
		Content string `form:"content" binding:"required"`
	}

	// Validate form data
	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Title and content are required",
		})
		return
	}

	// Retrieve the logged-in user from context
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User not authenticated",
		})
		return
	}

	// Create a new post
	post := models.Post{
		Title:   body.Title,
		Content: body.Content,
		UserID:  userID.(uint),
	}
	result := initializers.DB.Create(&post)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create post",
		})
		return
	}

	c.Redirect(http.StatusMovedPermanently, "/home")
	// Respond with the created post
	// c.JSON(http.StatusOK, gin.H{
	// 	"post": post,
	// })
}

// GetPost retrieves a post by ID
func GetPost(c *gin.Context) {
	// Get the post ID from the URL parameter
	postID := c.Param("id")

	// Find the post
	var post models.Post
	result := initializers.DB.Preload("User").First(&post, postID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Post not found",
		})
		return
	}

    var user models.User

	c.HTML(http.StatusOK, "viewblog.html", gin.H{
        "Post": post,
        "User": user,
    })
	// Respond with the post
	// c.JSON(http.StatusOK, gin.H{
	// 	"post": post,
	// })
}

// GetAllPosts retrieves all posts
// func GetAllPosts(c *gin.Context) {
// 	var posts []models.Post
// 	result := initializers.DB.Preload("User").Find(&posts)
// 	if result.Error != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": "Failed to retrieve posts",
// 		})
// 		return
// 	}

// 	// Render the posts template with the list of posts
// 	c.HTML(http.StatusOK, "home.html", gin.H{
// 		"Posts": posts,
// 	})
// 	// Respond with the list of posts
// 	// c.JSON(http.StatusOK, gin.H{
// 	// 	"posts": posts,
// 	// })
// }

// UpdatePost updates a post by ID
func UpdatePost(c *gin.Context) {
	// Get the post ID from the URL parameter
	postID := c.Param("id")

	// Parse form data
	var body struct {
		Title   string `form:"title"`
		Content string `form:"content"`
	}

	// Validate form data
	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input",
		})
		return
	}

	// Find the post
	var post models.Post
	result := initializers.DB.First(&post, postID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Post not found",
		})
		return
	}

	// Update the post fields
	if body.Title != "" {
		post.Title = body.Title
	}
	if body.Content != "" {
		post.Content = body.Content
	}
	initializers.DB.Save(&post)

	c.Redirect(http.StatusFound, "/posts/"+postID)
	// c.Redirect(http.StatusMovedPermanently, "/posts")
	// Respond with the updated post
	// c.JSON(http.StatusOK, gin.H{
	// 	"post": post,
	// })
}

func EditPost(c *gin.Context) {
    postID := c.Param("id")
	userID, exists := c.Get("user_id")
	fmt.Println(userID)

	var user models.User
	if exists {
		initializers.DB.First(&user, userID)
	}

    // if err := initializers.DB.First(&post, id).Error; err != nil {
    //     c.HTML(http.StatusNotFound, "error.html", gin.H{
    //         "error": "Post not found",
    //     })
    //     return
    // }
	var post models.Post

	result := initializers.DB.Preload("User").First(&post, postID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Post not found",
		})
		return
	}

	fmt.Println(post.UserID)
	
	

	// result := initializers.DB.Preload("User").First(&post, postID)
	// if result.Error != nil {
	// 	c.JSON(http.StatusNotFound, gin.H{
	// 		"error": "Post not found",
	// 	})
	// 	return
	// }

    // userID, _ := c.Get("user_id")
    if user.ID != post.UserID {
        c.HTML(http.StatusForbidden, "error.html", gin.H{
            "error": "Unauthorized to edit this post",
        })
        return
    }


    // var user models.User
    // initializers.DB.First(&user, userID.(uint))

    c.HTML(http.StatusOK, "editblog.html", gin.H{
        "Post": post,
        // "User": user,
    })
}
// DeletePost deletes a post by ID
func DeletePost(c *gin.Context) {
	// Get the post ID from the URL parameter
	postID := c.Param("id")

	// Find and delete the post
	result := initializers.DB.Delete(&models.Post{}, postID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Post not found",
		})
		return
	}

	// Respond with a success message
	c.JSON(http.StatusOK, gin.H{
		"message": "Post deleted successfully",
	})
}

func ShowCreatePostPage(c *gin.Context) {
	userID, exists := c.Get("user_id")
	var user models.User

	if exists {
		initializers.DB.First(&user, userID)
	}

	c.HTML(http.StatusOK, "addblog.html", gin.H{
		"User": user,
	})
}

// ShowPostsPage displays a list of posts
func ShowPostsPage(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Preload("User").Find(&posts)

	userID, exists := c.Get("user_id")
	var user models.User

	if exists {
		initializers.DB.First(&user, userID)
	}

	c.HTML(http.StatusOK, "home.html", gin.H{
		"Posts": posts,
		"User":  user,
	})
}