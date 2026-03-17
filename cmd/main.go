package main

import (
	"MackaWebsite/internal/database"
	"MackaWebsite/internal/handlers"
	"MackaWebsite/internal/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	if err := database.ConnectToDb(); err != nil {
		return
	}

	if err := database.InitializeDatatable(); err != nil {
		return
	}

	_ = database.CreatTheFirstUser()

	app.Use(middleware.LimitRequestBody(2 * 1024 * 1024))

	app.Group("/api")
	app.POST("/login", handlers.Login)
	app.GET("/landing-page", handlers.GetLandingPage)
	app.GET("/blogs", handlers.GetBlogs)
	app.GET("/blogs/:title", handlers.GetBlogByTitle)
	app.PUT("/blogs/:title", handlers.UpdateBlogByTitle)
	app.PATCH("/blogs/:title", handlers.UpdateBlogImageByTitle)
	app.POST("/blogs", handlers.PostBlog)

	app.Run(":8080")
}
