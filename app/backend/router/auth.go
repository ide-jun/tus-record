package router

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func addAuthRouter(rg *gin.RouterGroup, h *controllers.Handler) {
	auth := rg.Group("/auth")

	auth.POST("/sign-up", h.SignUpHandler)
	auth.POST("/sign-in", h.SignInHandler)
	auth.GET("/get-data", controllers.AuthMiddleware, h.GetUserData)
}
