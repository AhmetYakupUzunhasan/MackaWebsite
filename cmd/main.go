package main

import (
	"MackaWebsite/internal/database"
	"MackaWebsite/internal/handlers"

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

	app.POST("/login", handlers.Login)

	app.Run(":8080")
}
