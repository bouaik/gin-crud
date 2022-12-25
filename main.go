package main

import (
	"gin-grud/controllers"
	"gin-grud/initializers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.PostsGet)
	r.GET("/posts/:id", controllers.PostShow)
	r.PUT("/posts/:id", controllers.PostUpdate)
	r.DELETE("/posts/:id", controllers.PostDelete)
	r.Run()
}
