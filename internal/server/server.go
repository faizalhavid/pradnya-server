package server

import (
	"fmt"

	"github.com/faizalhavid/pradnya-server/internal/config"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Router *gin.Engine
	Config *config.Config
}

func New(cfg *config.Config) *Server {
	return &Server{
		Config: cfg,
		Router: setupRouter(),
	}
}

func (s *Server) Run() error {
	address := fmt.Sprintf(":%s", s.Config.App.AppPort)
	return s.Router.Run(address)
}
