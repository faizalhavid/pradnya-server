package auth

import (
	"time"

	"github.com/faizalhavid/pradnya-server/internal/shared"
	"github.com/faizalhavid/pradnya-server/internal/user"
)

type DeviceInfo struct {
	shared.BaseModel

	DeviceName  string    `json:"device_name"`
	DeviceOS    string    `json:"device_os"`
	LastLoginAt time.Time `json:"last_login"`
	IpAddress   string    `json:"ip_adddress"`
	UserAgent   string    `json:"user_agent"`
	Fingerprint string    `gorm:"uniqueIndex"`

	UserId string    `gorm:"index;not null" json:"user_id"`
	User   user.User `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE" json:"user"`
}
