package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/faizalhavid/pradnya-server/internal/shared"
	"github.com/gin-gonic/gin"
)

const ContextUserID = "user_id"

func AuthMiddleware(
	jwtCfg shared.JWTConfig,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		fmt.Println("authHeader", authHeader)
		fmt.Println("header", c.Request.Header)

		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "missing authorization"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := jwtCfg.ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			return
		}

		if claims.Type != shared.TokenPurposeAccess {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"message": "invalid token purpose",
				},
			)
			return
		}
		c.Set(
			ContextUserID,
			claims.UserID,
		)

		c.Next()
	}
}
