package controllers

import (
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(context *gin.Context) {
	_, ok := checkAuthHeader(context)
	if !ok {
		return
	}

	context.Next()
}
