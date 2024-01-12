package api

import (
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
)

type API struct {
	db *sql.DB
}

func (s *API) ConnectDB() {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		// エラーハンドリング
		return
	}
	c := mysql.Config{
		DBName:    "tus-record",
		User:      "amdin",
		Passwd:    "pass",
		Addr:      "db:3306",
		Net:       "tcp",
		ParseTime: true,
		Collation: "utf8mb4_unicode_ci",
		Loc:       jst,
	}
	db, err := sql.Open("mysql", c.FormatDSN())
	if err != nil {
		// エラーハンドリング
		return
	}
	s.db = db
}
