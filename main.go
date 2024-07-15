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
	r.GET("/post/:id", controllers.PostGetById)
	r.PUT("/post/:id", controllers.PostUpdateById)
	r.DELETE("/post/:id", controllers.PostDelete)

	r.POST("/post", controllers.PostCreate)

	r.Run()
}
