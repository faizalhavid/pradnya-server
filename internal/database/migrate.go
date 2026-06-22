package database

import (
	model "github.com/faizalhavid/pradnya-server/internal/user"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.User{},
	)
}
