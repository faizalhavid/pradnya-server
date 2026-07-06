package middleware

import (
	"errors"
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
		claims, err := jwtCfg.ValidateToken(
			tokenString,
			shared.TokenPurposeAccess,
		)
		if err != nil {

			switch {

			case errors.Is(err, shared.ErrTokenExpired):
				c.AbortWithStatusJSON(
					http.StatusUnauthorized,
					gin.H{"message": "token expired"},
				)

			case errors.Is(err, shared.ErrTokenWrongPurpose):
				c.AbortWithStatusJSON(
					http.StatusUnauthorized,
					gin.H{"message": "invalid token purpose"},
				)

			default:
				c.AbortWithStatusJSON(
					http.StatusUnauthorized,
					gin.H{"message": "invalid token"},
				)
			}

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
