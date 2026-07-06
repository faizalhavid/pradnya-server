package auth

import (
	model "github.com/faizalhavid/pradnya-server/internal/user"
	"gorm.io/gorm"
)

type Repository interface {
	Create(user *model.User) error
	FindByEmail(email string) (*model.User, error)
	FindByID(id string) (*model.User, error)
	UpdatePassword(userId string, newPassword string) error
	UserTokenRepository
	DeviceInfoRepository
}
type repository struct {
	db *gorm.DB
	UserTokenRepository
	DeviceInfoRepository
}

func NewRepository(
	db *gorm.DB,
) Repository {
	return &repository{
		db:                   db,
		UserTokenRepository:  NewUserTokenRepository(db),
		DeviceInfoRepository: NewDeviceInfoRepository(db),
	}
}

var _ Repository = (*repository)(nil)

func (r *repository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *repository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.db.Where("email=?", email).First(&user).Error
	return &user, err
}

func (r *repository) FindByID(id string) (*model.User, error) {
	var user model.User
	err := r.db.Where("id=?", id).First(&user).Error
	return &user, err
}

func (r *repository) UpdatePassword(userId string, newPassword string) error {
	return r.db.Model(&model.User{}).Where("id=?", userId).Update("password", newPassword).Error
}

func (r *repository) CreateUserToken(token *UserToken) error {
	return r.db.Create(token).Error
}

func (r *repository) FirsOrCreateDeviceInfo(device *DeviceInfo) (*DeviceInfo, error) {
	var existingDevice DeviceInfo
	err := r.db.Where("fingerprint = ?", device.Fingerprint).First(&existingDevice).Error
	if err == nil {
		return &existingDevice, nil
	}
	if err != gorm.ErrRecordNotFound {
		return nil, err
	}

	err = r.db.Create(device).Error
	if err != nil {
		return nil, err
	}

	return device, nil
}
