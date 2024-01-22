package router

import (
	"github.com/ide-jun/tus-record/src/controllers"

	"github.com/gin-gonic/gin"
)

func addPingRouter(rg *gin.RouterGroup, h *controllers.Handler) {
	ping := rg.Group("/ping")

	ping.POST("", controllers.AuthMiddleware, h.PingHandler)
}
