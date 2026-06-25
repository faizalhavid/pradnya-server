package server

import (
	"github.com/faizalhavid/pradnya-server/internal/auth"
	"github.com/faizalhavid/pradnya-server/internal/config"
	"github.com/faizalhavid/pradnya-server/internal/shared"
	"gorm.io/gorm"
)

type Modules struct {
	AuthHandler *auth.Handler
}

func BuildModules(
	db *gorm.DB,
	cfg *config.Config,
	mailer shared.Mailer,
) *Modules {
	authRepo := auth.NewRepository(db)

	authService := auth.NewService(
		authRepo,
		shared.JWTConfig{
			Secret: cfg.JWT.Secret,
			Issuer: cfg.App.AppName,
		},
		mailer,
	)
	authHandler := auth.NewHandler(
		authService,
	)

	return &Modules{
		AuthHandler: authHandler,
	}
}
