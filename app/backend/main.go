package main

import (
	"backend/config"
	"backend/database"
	"backend/router"
	"log"
	"time"
)

const (
	// corsMaxAge プリフライトリクエストをキャッシュできる時間 (秒)。
	corsMaxAge = 600
	// readTimeout リクエスト読み取りのタイムアウト。
	readTimeout = 5 * time.Second
	// writeTimeout レスポンス書き込みのタイムアウト。
	writeTimeout = 10 * time.Second
)

func main() {
	cfg, err := config.Get()
	if err != nil {
		log.Fatal(err)
	}

	if err := database.SetUpDB(cfg.DBHost, cfg.DBPort, cfg.DBName, cfg.DBUsername, cfg.DBPassword); err != nil {
		log.Fatal(err)
	}

	router.Run()
}
