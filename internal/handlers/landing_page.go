package handlers

import (
	"MackaWebsite/internal/database"
	"MackaWebsite/internal/models"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetLandingPage(ctx *gin.Context) {
	landingPage, err := database.SelectLandingPageFromDb()
	if err != nil {
		ctx.JSON(500, gin.H{
			"data":  nil,
			"error": err,
		})
		return
	}

	ctx.JSON(200, gin.H{
		"data":  landingPage,
		"error": nil,
	})
}

func GetBlogByTitle(ctx *gin.Context) {
	title := ctx.Param("title")
	blog, err := database.SelectBlogFromDbById(title)
	if err != nil {
		ctx.JSON(500, gin.H{
			"data":  nil,
			"error": err,
		})
		return
	}

	ctx.JSON(200, gin.H{
		"data":  blog,
		"error": nil,
	})
}

func GetBlogs(ctx *gin.Context) {
	blogs, err := database.SelectBlogsFromDb()
	if err != nil {
		ctx.JSON(500, gin.H{
			"data":  nil,
			"error": err,
		})
		return
	}

	ctx.JSON(200, gin.H{
		"data":  blogs,
		"error": nil,
	})
}

func UpdateBlogImageByTitle(ctx *gin.Context) {
	title := ctx.Param("title")
	file, err := ctx.FormFile("file")
	if err != nil {
		if strings.Contains(err.Error(), "http: request body too large") {
			ctx.JSON(413, gin.H{
				"data":  nil,
				"error": err.Error(),
			})
			return
		}

		fmt.Println("You Monster")
		ctx.JSON(500, gin.H{
			"data":  nil,
			"error": err,
		})
		return
	}

	ext := filepath.Ext(file.Filename)
	if ext != ".png" && ext != ".jpeg" && ext != ".jpg" {
		ctx.JSON(400, gin.H{
			"data":  nil,
			"error": "File Type is not suported",
		})
		return
	}

	if err := database.VerifyBlogFromDbByTitle(title); err != nil {
		ctx.JSON(400, gin.H{
			"data":  nil,
			"error": err.Error(),
		})
		return
	}

	dst := fmt.Sprintf("Uploads/%s", file.Filename)
	if err := ctx.SaveUploadedFile(file, dst); err != nil {
		ctx.JSON(500, gin.H{
			"data":  nil,
			"error": err,
		})
		return
	}

	if err := database.UpdateBlogImageLinkByTitleInDb(dst, title); err != nil {
		ctx.JSON(500, gin.H{
			"data":  nil,
			"error": err,
		})
		return
	}

	ctx.JSON(204, gin.H{
		"data":  nil,
		"error": nil,
	})
}

func PostBlog(ctx *gin.Context) {
	var blog models.Blog
	if err := ctx.ShouldBindJSON(&blog); err != nil {
		ctx.JSON(400, gin.H{
			"data":  nil,
			"error": err.Error(),
		})
		return
	}

	if err := database.InsertBlogIntoDb(&blog); err != nil {
		ctx.JSON(500, gin.H{
			"data":  nil,
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(201, gin.H{
		"data":  blog,
		"error": nil,
	})
}

func UpdateBlogByTitle(ctx *gin.Context) {
	title := ctx.Param("title")

	var blog models.Blog
	if err := ctx.ShouldBindJSON(&blog); err != nil {
		ctx.JSON(400, gin.H{
			"data":  nil,
			"error": err.Error(),
		})
		return
	}

	if err := database.UpdateBlogByTitleInDb(title, &blog); err != nil {
		ctx.JSON(400, gin.H{
			"data":  nil,
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"data":  blog,
		"error": nil,
	})
}
