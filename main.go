package main

import (
	"github.com/YerkinSuimenbay/go-crud-posts/controllers"
	"github.com/YerkinSuimenbay/go-crud-posts/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.PostsGet)
	r.GET("/posts/:id", controllers.PostsGetById)
	r.PUT("/posts/:id", controllers.PostsUpdate)
	r.DELETE("/posts/:id", controllers.PostsDelete)

	r.Run()
}