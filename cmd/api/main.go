package main

import (
	"log"

	"github.com/faizalhavid/pradnya-server/internal/config"
	"github.com/faizalhavid/pradnya-server/internal/database"
	"github.com/faizalhavid/pradnya-server/internal/server"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db, err := database.ConnectPostgres(cfg.DB)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	if err := database.Migrate(db); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	//
	srv := server.New(cfg)
	if err := srv.Run(); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
