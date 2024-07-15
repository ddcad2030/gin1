package main

import (
	"example/gin-crud-1/controllers"
	"example/gin-crud-1/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {
	r := gin.Default()

	r.GET("/post", controllers.PostGet)
	r.POST("/post", controllers.PostCreate)

	r.Run()
}
