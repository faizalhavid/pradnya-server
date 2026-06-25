package server

import (
	"github.com/faizalhavid/pradnya-server/internal/auth"
	"github.com/faizalhavid/pradnya-server/internal/config"
	"github.com/faizalhavid/pradnya-server/internal/middleware"
	"github.com/faizalhavid/pradnya-server/internal/shared"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Modules struct {
	AuthHandler    *auth.Handler
	AuthMiddleware gin.HandlerFunc
}

func BuildModules(
	db *gorm.DB,
	cfg *config.Config,
	mailer shared.Mailer,
) *Modules {
	authRepo := auth.NewRepository(db)
	jwtCfg := shared.JWTConfig{
		Secret: cfg.JWT.Secret,
		Issuer: cfg.App.AppName,
	}
	authService := auth.NewService(
		authRepo,
		jwtCfg,
		mailer,
	)
	authMiddleware := middleware.AuthMiddleware(
		jwtCfg,
	)
	authHandler := auth.NewHandler(
		authService,
	)

	return &Modules{
		AuthHandler:    authHandler,
		AuthMiddleware: authMiddleware,
	}
}
