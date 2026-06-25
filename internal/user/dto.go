package user

import "time"

type ProfileRequest struct {
	FirstName string    `json:"first_name" binding:"required" example:"John"`
	LastName  string    `json:"last_name" binding:"required" example:"Doe"`
	DateBirth time.Time `json:"date_birth" binding:"required" example:"1990-01-01"`
	Photo     string    `json:"photo" binding:"required" example:"https://example.com/photo.jpg"`
	Gender    string    `json:"gender" binding:"required" example:"Male"`
}

type UserResponse struct {
	ID       string          `json:"id" example:"1"`
	Username string          `json:"username" example:"johndoe"`
	Email    string          `json:"email" example:"johndoe@example.com"`
	Profile  ProfileResponse `json:"profile"`
}

type ProfileResponse struct {
	FirstName string    `json:"first_name" example:"John"`
	LastName  string    `json:"last_name" example:"Doe"`
	DateBirth time.Time `json:"date_birth" example:"1990-01-01"`
	Photo     string    `json:"photo" example:"https://example.com/photo.jpg"`
	Gender    string    `json:"gender" example:"Male"`
}
