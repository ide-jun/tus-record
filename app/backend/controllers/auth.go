package controllers

import (
	"backend/clock"
	"backend/models"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type Handler struct {
	DB  *gorm.DB
	Clk clock.Clock
}

func (handler *Handler) SignInHandler(context *gin.Context) {
	var signInInput models.SignInInput
	if err := context.ShouldBind(&signInInput); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	user := &models.User{
		Email:    signInInput.Email,
		Password: signInInput.Password,
	}

	if err := user.CheckUser(handler.DB); err != nil {
		if err.Code == http.StatusUnauthorized {
			context.JSON(err.Code, gin.H{
				"message": err.Message,
			})
		}
	}

	// token の発行
	token := jwt.New(jwt.GetSigningMethod("HS256"))

	// claims の設定
	token.Claims = jwt.MapClaims{
		"user": fmt.Sprint(user.ID),
		"exp":  handler.Clk.Now().Add(time.Hour * 1).Unix(),
	}

	// 署名
	secretKey := "secret"
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to create token!",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"user_id": user.ID,
		"message": "Successfully sign in",
		"token":   tokenString,
	})
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

	if err := newUser.Create(handler.DB); err != nil {
		if err.Code == http.StatusConflict {
			context.JSON(err.Code, gin.H{
				"message": err.Message,
			})
		}
		if err.Code == http.StatusBadRequest {
			context.JSON(err.Code, gin.H{
				"message": err.Message,
			})
		}
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"user_id": newUser.ID,
		"message": "Successfully created user",
	})
}
