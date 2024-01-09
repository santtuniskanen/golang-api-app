package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/santtuniskanen/golang-restapi/database"
	"github.com/santtuniskanen/golang-restapi/models"
)

func CreatePost(c *gin.Context) {

	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	post := models.Post{Title: body.Title, Body: body.Body}

	result := database.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func GetPosts(c *gin.Context) {
	var posts []models.Post
	database.DB.Find(&posts)

	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

func GetPostByID(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	database.DB.First(&post, id)

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func UpdatePost(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	var post models.Post
	database.DB.First(&post, id)

	database.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	database.DB.First(&post, id)

	database.DB.Unscoped().Delete(&post, id)

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}
