package app

import (
	"backend/clock"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

// API API を表す構造体
type API struct {
	// clk 現在時刻を取得するインターフェース
	clk clock.Clock
	// db データベースハンドラ
	db *sql.DB

	// middlewares API で使用する HTTP ミドルウェア
	middlewares []func(http.Handler) http.Handler
}

func NewAPI(clk clock.Clock, db *sql.DB) *API {
	return &API{
		clk:         clk,
		db:          db,
		middlewares: make([]func(http.Handler) http.Handler, 0),
	}
}

func (s *API) Handler() http.Handler {
	r := chi.NewRouter()

	r.Use(s.middlewares...)

	r.Post("/ping", s.pingHandler)

	return r
}

type PingRequest struct {
	Message string `json:"message"`
}

type PingResponse struct {
	Message      string    `json:"message"`
	ReceivedTime time.Time `json:"receivedTime"`
}

func (s *API) pingHandler(w http.ResponseWriter, r *http.Request) {
	receivedTime := s.clk.Now()

	// ReqBを デコード
	req := &PingRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		// デコードに失敗した場合はログ出力して 400 Bad Request を返す。
		log.Printf("ERROR: request decoding failed! %+v", err)
		writeErrorResponse(w, http.StatusBadRequest, "invalid json")
		return
	}

	// RespB の構造体を生成
	resp := &PingResponse{
		Message:      req.Message,
		ReceivedTime: receivedTime,
	}

	// RespB をエンコード
	if err := json.NewEncoder(w).Encode(&resp); err != nil {
		// エンコードに失敗した場合はログ出力して 500 Internal Server Error を返す。
		log.Printf("ERROR: response encoding failed: %+v", err)
		writeHTTPError(w, http.StatusInternalServerError)
		return
	}
}
