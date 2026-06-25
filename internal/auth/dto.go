package auth

import (
	"github.com/faizalhavid/pradnya-server/internal/shared"
	"github.com/faizalhavid/pradnya-server/internal/user"
)

type RegisterRequest struct {
	Name     string `json:"name" binding:"required" example:"John Doe"`
	Email    string `json:"email" binding:"required,email" example:"johndoe@example.com"`
	Password string `json:"password" binding:"required,min=8" example:"password123"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email" example:"johndoe@example.com"`
	Password string `json:"password" binding:"required" example:"password123"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email" binding:"required,email" example:"johndoe@example.com"`
}

type ResetPasswordRequest struct {
	Token       string `json:"token" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=8" example:"newpassword123"`
}

type CredentialsData struct {
	AccessToken  shared.TokenData `json:"access_token"`
	RefreshToken shared.TokenData `json:"refresh_token"`
}

type LoginResponse struct {
	User        user.UserResponse `json:"user"`
	Credentials CredentialsData   `json:"credentials"`
}

type RegisterResponse struct {
	User user.UserResponse `json:"user"`
}

type ForgotPasswordResponse struct {
	ToEmail    string           `json:"to_email"`
	ResetToken shared.TokenData `json:"reset_token"`
}
