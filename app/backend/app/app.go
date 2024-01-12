package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type App struct {
	api *API

	// middlewares サーバーアプリケーション全体で使用する HTTP ミドルウェア
	middlewares []func(http.Handler) http.Handler
}

func New(api *API) *App {
	app := &App{
		api:         api,
		middlewares: make([]func(http.Handler) http.Handler, 0),
	}

	return app
}

func (s *App) Handler() http.Handler {
	r := chi.NewRouter()

	r.Use(s.middlewares...)
	r.Use(middleware.NoCache)

	r.Mount("/", s.api.Handler())

	return r
}
