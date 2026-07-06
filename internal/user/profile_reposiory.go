package user

import (
	"gorm.io/gorm"
)

type ProfileRepository interface {
	Create(profile *UserProfile) error
	FindByUserID(userId string) (*UserProfile, error)
	UpdateByUserID(userId string, profile *UserProfile) error
	SoftDeleteByUserID(userId string) error
	PermanentlyDeleteByUserID(userId string) error
}
type profileRepository struct {
	db *gorm.DB
}

var _ ProfileRepository = (*profileRepository)(nil)

func NewProfileRepository(
	db *gorm.DB,
) ProfileRepository {
	return &profileRepository{
		db: db,
	}
}

func (r *profileRepository) Create(profile *UserProfile) error {
	return r.db.Create(profile).Error
}

func (r *profileRepository) FindByUserID(userId string) (*UserProfile, error) {
	var profile UserProfile
	err := r.db.Where("user_id=?", userId).First(&profile).Error
	return &profile, err
}

func (r *profileRepository) UpdateByUserID(userId string, profile *UserProfile) error {
	return r.db.Model(&UserProfile{}).Where("user_id=?", userId).Updates(profile).Error
}

func (r *profileRepository) SoftDeleteByUserID(userId string) error {
	return r.db.Model(&UserProfile{}).Where("user_id=?", userId).Update("is_deleted", true).Error
}

func (r *profileRepository) PermanentlyDeleteByUserID(userId string) error {
	return r.db.Model(&UserProfile{}).Where("user_id=?", userId).Delete(&UserProfile{}).Error
}
