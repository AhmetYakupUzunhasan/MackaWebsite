package handlers

import (
	"MackaWebsite/internal/database"
	"MackaWebsite/internal/middleware"
	"MackaWebsite/internal/models"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	var admin models.Admin
	err := ctx.ShouldBindJSON(&admin)
	if err != nil {
		ctx.JSON(400, gin.H{
			"data":  nil,
			"error": err.Error(),
		})
		return
	}

	lenUsername := utf8.RuneCountInString(admin.Username)
	if lenUsername < 5 {
		ctx.JSON(400, gin.H{
			"data":  nil,
			"error": "Username must be no less than 5 chars",
		})
		return
	}

	lenPassword := utf8.RuneCountInString(admin.Password)
	if lenPassword < 8 {
		ctx.JSON(400, gin.H{
			"data":  nil,
			"error": "Password legth must be greater than or equal to 8",
		})
		return
	}

	id, password, err := database.SelectUserPasswordByUsername(admin.Username)
	if err != nil {
		ctx.JSON(500, gin.H{
			"data":  nil,
			"error": err.Error(),
		})
		return
	}

	if password != admin.Password {
		ctx.JSON(401, gin.H{
			"data":  nil,
			"error": "Password or Username is Incorrect",
		})
		return
	}

	token, err := middleware.GenerateToken(id, "admin")
	if err != nil {
		ctx.JSON(500, gin.H{
			"data":  nil,
			"error": err,
		})
		return
	}

	ctx.JSON(200, gin.H{
		"data":  token,
		"error": nil,
	})
}
