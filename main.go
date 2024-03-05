package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-crud/initializers"
	"github.com/go-crud/controllers"
	"github.com/gin-contrib/cors"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}


func main() {
	
	r := gin.Default()
	r.Use(cors.Default())

	r.POST("/github/callback", controllers.HandleGitHubCallback)
	r.POST("/gitlab/callback", controllers.HandleGitLabCallback)

	r.POST("/posts", controllers.PostsCreate)
	r.PUT("/posts/:id", controllers.PostsUpdate)
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostsShow)
	r.DELETE("/posts/:id", controllers.PostsDelete)

	r.Run() // listen and serve on 0.0.0.0:8080
	fmt.Println("Hello, World!")
}