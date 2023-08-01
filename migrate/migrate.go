package main

import (
	"github.com/YerkinSuimenbay/go-crud-posts/initializers"
	"github.com/YerkinSuimenbay/go-crud-posts/models"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}