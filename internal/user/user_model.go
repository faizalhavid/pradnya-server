package user

import (
	"github.com/faizalhavid/pradnya-server/internal/shared"
)

type User struct {
	shared.BaseModel

	Username string      `gorm:"uniqueIndex;not null" json:"username"`
	Email    string      `gorm:"uniqueIndex;not null" json:"email"`
	Password string      `gorm:"not null" json:"-"`
	Profile  UserProfile `gorm:"constraint:OnDelete:CASCADE" json:"profile"`
}
