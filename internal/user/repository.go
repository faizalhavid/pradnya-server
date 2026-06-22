package user

import (
	"gorm.io/gorm"
)

type Repository interface {
	CreateProfile(user *UserProfile) error
	UpdateProfile(userID string, profile *UserProfile) error
	Delete(userId string) error
	PermanentlyDelete(userId *string) error
}
type repository struct {
	db *gorm.DB
}

var _ Repository = (*repository)(nil)

func NewRepository(
	db *gorm.DB,
) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateProfile(user *UserProfile) error {
	return r.db.Create(user).Error
}

func (r *repository) UpdateProfile(userId string, profile *UserProfile) error {
	return r.db.Model(&UserProfile{}).Where("userID=?", userId).Updates(profile).Error
}

func (r *repository) PermanentlyDelete(userId *string) error {
	return r.db.Model(&User{}).Where("userID=?", *userId).Delete(&User{}).Error
}

func (r *repository) Delete(userId string) error {
	return r.db.Model(&User{}).Where("userID=?", userId).Update("is_deleted", true).Error
}
