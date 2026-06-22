package main

import (
	"log"

	"github.com/faizalhavid/pradnya-server/internal/config"
	"github.com/faizalhavid/pradnya-server/internal/database"
	"github.com/faizalhavid/pradnya-server/internal/database/seed"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db, err := database.ConnectPostgres(cfg.DB)

	if err := seed.Run(db); err != nil {
		log.Fatal(err)
	}

	log.Println("seed completed")
}
