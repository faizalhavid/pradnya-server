package user

import (
	"time"

	"github.com/faizalhavid/pradnya-server/internal/shared"
)

type UserProfile struct {
	shared.BaseModel

	FirstName string    `gorm:"type:varchar(25);not null" json:"first_name"`
	LastName  string    `gorm:"type:varchar(25);not null" json:"last_name"`
	DateBirth time.Time `gorm:"not null" json:"date_birth"`
	Photo     string    `json:"photo_profile"`
	Gender    string    `gorm:"type:varchar(10);not null"`

	UserId string `json:"user_id"`
}
