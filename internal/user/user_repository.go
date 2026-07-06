package user

import (
	"github.com/faizalhavid/pradnya-server/internal/shared"
	"gorm.io/gorm"
)

type UserRepository interface {
	shared.BaseRepository[User]
}
type userRepository struct {
	shared.BaseRepository[User]
	db *gorm.DB
}

var _ UserRepository = (*userRepository)(nil)

func NewUserRepository(
	db *gorm.DB,
) UserRepository {
	return &userRepository{
		db: db,
	}
}
