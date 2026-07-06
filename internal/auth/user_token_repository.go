package auth

import (
	"gorm.io/gorm"
)

type UserTokenRepository interface {
	CreateUserToken(token *UserToken) error

	FindUserTokenByHash(hash string) (*UserToken, error)

	UpdateUserToken(token *UserToken) error

	RevokeUserToken(id string) error

	RevokeAllUserToken(userID string) error
}
type userTokenRepository struct {
	db *gorm.DB
}

func NewUserTokenRepository(
	db *gorm.DB,
) UserTokenRepository {
	return &userTokenRepository{
		db: db,
	}
}

var _ UserTokenRepository = (*userTokenRepository)(nil)

func (r *userTokenRepository) CreateUserToken(token *UserToken) error {
	return r.db.Create(token).Error
}

func (r *userTokenRepository) FindUserTokenByHash(hash string) (*UserToken, error) {
	var token UserToken
	err := r.db.Where("token_hash = ?", hash).First(&token).Error
	return &token, err
}

func (r *userTokenRepository) UpdateUserToken(token *UserToken) error {
	return r.db.Model(&token).Where("id = ?", token.ID).Updates(token).Error
}

func (r *userTokenRepository) RevokeUserToken(id string) error {
	return r.db.Model(&UserToken{}).Where("id = ?", id).Update("revoked", true).Error
}

func (r *userTokenRepository) RevokeAllUserToken(userID string) error {
	return r.db.Model(&UserToken{}).Where("user_id = ?", userID).Update("revoked", true).Error
}
