package seed

import (
	model "github.com/faizalhavid/pradnya-server/internal/user"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserSeeder struct{}

func (s UserSeeder) Run(db *gorm.DB) error {
	var count int64

	db.Model(&model.User{}).Where(
		"email = ?",
		"admin@example.com",
	).Count(&count)

	if count > 0 {
		return nil
	}

	password, err := bcrypt.GenerateFromPassword(
		[]byte("admin123"),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	admin := model.User{
		Username: "admin",
		Email:    "admin@example.com",
		Password: string(password),
	}

	return db.Create(&admin).Error
}
