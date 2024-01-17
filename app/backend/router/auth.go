package router

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func addAuthRouter(rg *gin.RouterGroup, h *controllers.Handler) {
	auth := rg.Group("/auth")

	auth.POST("/sign-up", h.SignUpHandler)
}
