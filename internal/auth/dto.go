package auth

import (
	"github.com/faizalhavid/pradnya-server/internal/shared"
	"github.com/faizalhavid/pradnya-server/internal/user"
)

type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type CredentialsData struct {
	AccessToken  shared.TokenData `json:"access_token"`
	RefreshToken shared.TokenData `json:"refresh_token"`
}

type LoginResponse struct {
	user        user.UserResponse
	credentials CredentialsData
}

type RegisterResponse struct {
	user user.UserResponse
}
