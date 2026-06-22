package user

import "time"

type ProfileRequest struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	DateBirth time.Time `json:"date_birth"`
	Photo     string    `json:"photo"`
	Gender    string    `json:"gender"`
}

type UserResponse struct {
	ID       string          `json:"id"`
	Username string          `json:"username"`
	Email    string          `json:"email"`
	Profile  ProfileResponse `json:"profile"`
}

type ProfileResponse struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	DateBirth time.Time `json:"date_birth"`
	Photo     string    `json:"photo"`
	Gender    string    `json:"gender"`
}
