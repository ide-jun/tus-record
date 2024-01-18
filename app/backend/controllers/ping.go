package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *Handler) PingHandler(context *gin.Context) {
	receivedTime := handler.Clk.Now()
	context.JSON(http.StatusOK, gin.H{
		"receivedTime": receivedTime,
	})
}
