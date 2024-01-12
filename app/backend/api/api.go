package api

import (
	"backend/config"
	"backend/database"
	"database/sql"
	"log"
)

type API struct {
	db *sql.DB
}

func (s *API) ConnectDB() {
	cfg, err := config.Get()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.Open(cfg.DBHost, cfg.DBPort, cfg.DBName, cfg.DBUsername, cfg.DBPassword)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	s.db = db
}
