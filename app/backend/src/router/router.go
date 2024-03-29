package router

import (
	"time"

	"github.com/ide-jun/tus-record/src/clock"
	"github.com/ide-jun/tus-record/src/controllers"
	"github.com/ide-jun/tus-record/src/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Run() {
	router := setupRouter()
	router.Run()
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	setupCORS(router)
	handler := &controllers.Handler{
		DB:  database.GetDB(),
		Clk: clock.New(),
	}

	api := router.Group("/api")
	addAuthRouter(api, handler)
	addPingRouter(api, handler)
	// addRecordRouter(api, handler)

	return router
}

func setupCORS(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		// アクセスを許可したいアクセス元
		AllowOrigins: []string{
			"http://localhost:5173",
		},
		// アクセスを許可したいHTTPメソッド
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
		},
		// 許可したいHTTPリクエストヘッダ
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		// cookieなどの情報を必要とするかどうか
		AllowCredentials: true,
		// preflightリクエストの結果をキャッシュする時間
		MaxAge: 24 * time.Hour,
	}))
}
