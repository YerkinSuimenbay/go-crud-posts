package controllers

import (
	"net/http"

	"github.com/YerkinSuimenbay/go-crud-posts/initializers"
	"github.com/YerkinSuimenbay/go-crud-posts/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(ctx *gin.Context) {
	// get data from req.body
	var body struct {
		Title string
		Body string
	}

	if err := ctx.Bind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()},
		)
		return
	}

	// create a post
	post := models.Post{Title: body.Title, Body: body.Body}
	
	result := initializers.DB.Create(&post)
	if result.Error != nil {
		ctx.Status(400)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func PostsGet(ctx *gin.Context) {
	var posts []models.Post
	result := initializers.DB.Find(&posts)

	if result.Error != nil {
		ctx.Status(400)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"count": result.RowsAffected,
		"posts": posts,
	})
}

func PostsGetById(ctx *gin.Context) {
	id := ctx.Param("id")

	var post models.Post

	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		ctx.Status(404)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func PostsUpdate(ctx *gin.Context) {
	id := ctx.Param("id")

	var body struct {
		Title string
		Body string
	}

	if err := ctx.Bind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()},
		)
		return
	}

	var post models.Post
	initializers.DB.First(&post, id)
	
	// if body.Title != "" {
	// 	post.Title = body.Title
	// }
	// if body.Body != "" {
	// 	post.Body = body.Body
	// }

	// initializers.DB.Save(&post)

	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})

	ctx.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func PostsDelete(ctx *gin.Context) {
	id := ctx.Param("id")

	initializers.DB.Delete(&models.Post{}, id)

	ctx.Status(200)
}