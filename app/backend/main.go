package main

import (
	"backend/api"
	"backend/config"
	"backend/database"
	"log"
)

type APP struct {
	api *api.API
}

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
}
