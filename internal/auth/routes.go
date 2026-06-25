package auth

import "github.com/gin-gonic/gin"

func RegisterRoutes(
	router *gin.RouterGroup,
	handler *Handler,
	authMiddleware gin.HandlerFunc,
) {
	auth := router.Group("/auth")

	auth.POST("/register", handler.Register)
	auth.POST("/login", handler.Login)
	auth.POST("/forgot-password", handler.ForgotPassword)
	auth.POST("/reset-password", handler.ResetPassword)
	protected := auth.Group("")
	protected.Use(authMiddleware)
	protected.GET(
		"/me",
		handler.Me,
	)
}
