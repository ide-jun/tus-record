package main

import (
	"backend/app"
	"backend/clock"
	"backend/config"
	"backend/database"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
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

	db, err := database.Open(cfg.DBHost, cfg.DBPort, cfg.DBName, cfg.DBUsername, cfg.DBPassword)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	clk := clock.New()

	api := app.NewAPI(clk, db)

	srvApp := app.New(api)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "DELETE"},
		AllowCredentials: true,
	})

	r.Mount("/", c.Handler(srvApp.Handler()))

	srv := &http.Server{
		Handler:      r,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}
	defer srv.Close()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	port := listener.Addr().(*net.TCPAddr).Port

	log.Printf("Server listening on port %d!", port)
	if err := srv.Serve(listener); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}
