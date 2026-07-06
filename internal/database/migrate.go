package database

import (
	"github.com/faizalhavid/pradnya-server/internal/auth"
	"github.com/faizalhavid/pradnya-server/internal/user"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&user.User{},
		&user.UserProfile{},

		&auth.UserToken{},
		&auth.DeviceInfo{},
	)
}
