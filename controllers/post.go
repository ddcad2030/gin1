package controllers

import (
	"example/gin-crud-1/initializers"
	"example/gin-crud-1/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {
	post := models.Post{Title: "Create", Body: "David"}
	result := initializers.DB.Create(&post)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, "Not create")
	}

	c.JSON(http.StatusCreated, post)

}

func PostGet(c *gin.Context) {
	var post models.Post
	result := initializers.DB.Find(&post)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, "Not found data")
	}
	c.JSON(http.StatusOK, result)
}
