package auth

import (
	"time"

	UserDto "github.com/faizalhavid/pradnya-server/internal/user"
)

type TokenPurpose string

const (
	TokenPurposeAccess  TokenPurpose = "access"
	TokenPurposeRefresh TokenPurpose = "refresh"
	TokenPurposeReset   TokenPurpose = "reset_password"
	TokenPurposeVerify  TokenPurpose = "verify_email"
)

type TokenData struct {
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expired_at"`
}

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
	AccessToken  TokenData `json:"access_token"`
	RefreshToken TokenData `json:"refresh_token"`
}

type LoginResponse struct {
	UserDto.UserResponse
	CredentialsData
}

type RegisterResponse struct {
	UserDto.UserResponse
}
