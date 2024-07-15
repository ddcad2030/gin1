package controllers

import (
	"example/gin-crud-1/initializers"
	"example/gin-crud-1/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {
	var bodyPost struct {
		Title string
		Body  string
	}
	c.Bind(&bodyPost)

	post := models.Post{Title: bodyPost.Title, Body: bodyPost.Body}
	result := initializers.DB.Create(&post)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, "Not create")
	}

	c.JSON(http.StatusCreated, post)

}

func PostGet(c *gin.Context) {
	var post []models.Post
	result := initializers.DB.Find(&post)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not found data",
		})
		return
	}
	c.JSON(http.StatusOK, post)
}

func PostGetById(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not found data",
		})
		return
	}
	c.JSON(http.StatusOK, post)
}

func PostUpdateById(c *gin.Context) {
	id := c.Param("id")
	var post models.Post

	result := initializers.DB.First(&post, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not found data",
		})
		return
	}

	var bodyPost struct {
		Title string
		Body  string
	}
	c.Bind(&bodyPost)

	resultUpdate := initializers.DB.Model(&post).Updates(models.Post{Title: bodyPost.Title, Body: bodyPost.Body})

	if resultUpdate.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Can not Update",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Update Success",
		"data":    post,
	})
}

func PostDelete(c *gin.Context) {
	id := c.Param("id")
	var post models.Post

	result := initializers.DB.First(&post, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not found data",
		})
		return
	}

	initializers.DB.Delete(&post, id)

	c.JSON(http.StatusOK, gin.H{
		"message": "Delete Successful",
	})
}
