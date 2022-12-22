package main

import (
	"gin-grud/controllers"
	"gin-grud/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.PostsGet)
	r.GET("/posts/:id", controllers.PostShow)
	r.Run()
}
