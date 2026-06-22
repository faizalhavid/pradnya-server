package user

import (
	"time"

	"github.com/faizalhavid/pradnya-server/internal/shared"
)

type User struct {
	shared.BaseModel

	Username string      `gorm:"uniqueIndex;not null" json:"username"`
	Email    string      `gorm:"uniqueIndex;not null" json:"email"`
	Password string      `gorm:"not null" json:"-"`
	Profile  UserProfile `gorm:"constraint:OnDelete:CASCADE" json:"profile"`
}

type UserProfile struct {
	shared.BaseModel

	FirstName string    `gorm:"type:varchar(25);not null" json:"first_name"`
	LastName  string    `gorm:"type:varchar(25);not null" json:"last_name"`
	DateBirth time.Time `gorm:"not null" json:"date_birth"`
	Photo     string    `json:"photo_profile"`
	Gender    string    `gorm:"type:varchar(10);not null"`

	UserId string `json:"user_id"`
}
