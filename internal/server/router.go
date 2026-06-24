package server

import (
	"github.com/faizalhavid/pradnya-server/internal/auth"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func setupRouter(modules *Modules) *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")

	api.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})
	api.GET(
		"/swagger/*any",
		ginSwagger.WrapHandler(
			swaggerFiles.Handler,
		),
	)

	auth.RegisterRoutes(
		api,
		modules.AuthHandler,
	)

	return router
}
