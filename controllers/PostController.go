package controllers

import (
	"net/http"

	"github.com/Smartdevs17/go_crud/initializers"
	"github.com/Smartdevs17/go_crud/models"
	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	//Get data from request body
	var body struct {
		Body   string
		Title  string
		Author string
	}

	c.Bind(&body)

	//Insert data into database
	newPost := models.Post{Title: body.Title, Body: body.Body, Author: body.Author}

	result := initializers.DB.Create(&newPost)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "error creating post",
		})
		return
	}

	//Return reponse
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "post successfully created",
		"data":    newPost,
	})
}

func GetAllPost(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)

	//return response
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "posts successfully retrieved",
		"data":    posts,
	})
}

func GetPost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post

	initializers.DB.First(&post, id)

	//return response
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "post successfully retrieved",
		"data":    post,
	})
}

func UpdatePost(c *gin.Context) {
	id := c.Param("id")
	var body struct {
		Body   string
		Title  string
		Author string
	}
	c.Bind(&body)
	var post models.Post
	initializers.DB.First(&post, id)
	post.Title = body.Title
	post.Author = body.Author
	post.Body = body.Body
	//save updated posts
	initializers.DB.Save(&post)

	//return response
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "post successfully updated",
		"data":    post,
	})
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	initializers.DB.Delete(&post, id)

	//return response
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "post successfully deleted",
		"data":    nil,
	})
}
