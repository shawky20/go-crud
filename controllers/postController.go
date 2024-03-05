package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-crud/initializers"
	"github.com/go-crud/models"
)

func PostsCreate (c *gin.Context) {
	// get the body of our POST request
	var requestBody struct {
		Body string `json:"body" binding:"required"`
		Title string `json:"title" binding:"required"`
	}
	
	c.Bind(&requestBody)

	// create post
	post := models.Post{Title: requestBody.Title, Body: requestBody.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"post": post,
		})
		return
	}
	// return the new post that we created
	c.JSON(200, gin.H{
		"post": post,
	})
}

// find all posts
func PostsIndex (c *gin.Context) {
	var posts []models.Post
	result := initializers.DB.Find(&posts)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"posts": posts,
		})
		return
	}

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

// find a post
func PostsShow (c *gin.Context) {
	var post models.Post
	result := initializers.DB.First(&post, c.Param("id"))

	if result.Error != nil {
		c.JSON(400, gin.H{
			"post": post,
		})
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

// update a post
func PostsUpdate (c *gin.Context) {
	// find a post by id param
	var post models.Post
	result := initializers.DB.First(&post, c.Param("id"))

	if result.Error != nil {
		c.JSON(400, gin.H{
			"post": post,
		})
		return
	}

	var requestBody struct {
		Body string `json:"body" binding:"required"`
		Title string `json:"title" binding:"required"`
	}
	c.Bind(&requestBody)

	post.Title = requestBody.Title
	post.Body = requestBody.Body

	result = initializers.DB.Save(&post)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"post": post,
		})
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

// delete a post
func PostsDelete (c *gin.Context) {
	var post models.Post
	result := initializers.DB.First(&post, c.Param("id"))

	if result.Error != nil {
		c.JSON(400, gin.H{
			"post": post,
		})
		return
	}

	result = initializers.DB.Delete(&post)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"post": post,
		})
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}