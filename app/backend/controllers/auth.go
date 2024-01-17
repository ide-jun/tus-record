package controllers

import (
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func (handler *Handler) SignUpHandler(context *gin.Context) {
	var signUpInput models.SignUpInput
	if err := context.ShouldBind(&signUpInput); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	newUser := &models.User{
		Name:     signUpInput.Name,
		Email:    signUpInput.Email,
		Password: signUpInput.Password,
	}

	if err := newUser.Validate(); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err := newUser.Create(handler.DB)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to create user",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"user_id": newUser.ID,
		"message": "Successfully created user",
	})
}
