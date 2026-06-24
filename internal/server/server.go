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

func NewServer(cfg *config.Config, modules *Modules) *Server {
	return &Server{
		Config: cfg,
		Router: setupRouter(modules),
	}
}

func (s *Server) Run() error {
	address := fmt.Sprintf(":%s", s.Config.App.AppPort)
	return s.Router.Run(address)
}
