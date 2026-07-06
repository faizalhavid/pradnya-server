package auth

import (
	"time"

	"github.com/faizalhavid/pradnya-server/internal/shared"
	"github.com/faizalhavid/pradnya-server/internal/user"
)

type UserToken struct {
	shared.BaseModel

	Purpose   shared.TokenPurpose `gorm:"index;not null" json:"purpose"`
	TokenHash string              `gorm:"type:varchar(64);uniqueIndex"`
	ExpiredAt time.Time           `gorm:"not null" json:"expired_at"`
	UsedAt    *time.Time          `json:"used_at"`
	RevokedAt *time.Time          `json:"revoked_at"`

	UserId   string     `gorm:"index;not null" json:"user_id"`
	User     user.User  `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE" json:"user"`
	DeviceID string     `gorm:"index"`
	Device   DeviceInfo `gorm:"foreignKey:DeviceID"`
}
