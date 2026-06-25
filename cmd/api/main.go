// @title Pradnya API
// @version 1.0
// @description Pradnya Backend API
// @BasePath /api

package main

import (
	"log"

	_ "github.com/faizalhavid/pradnya-server/docs"

	"github.com/faizalhavid/pradnya-server/internal/config"
	"github.com/faizalhavid/pradnya-server/internal/database"
	"github.com/faizalhavid/pradnya-server/internal/server"
	"github.com/faizalhavid/pradnya-server/internal/shared"
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

	mailer := shared.NewMailer(
		cfg.Mail.Host,
		cfg.Mail.Port,
		cfg.Mail.Username,
		cfg.Mail.Password,
		cfg.Mail.FromName,
	)
	modules := server.BuildModules(db, cfg, *mailer)
	srv := server.NewServer(cfg, modules)
	if err := srv.Run(); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
