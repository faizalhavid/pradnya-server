package auth

import (
	"gorm.io/gorm"
)

type DeviceInfoRepository interface {
	FirstOrCreateDeviceInfo(device *DeviceInfo) (*DeviceInfo, error)
}
type deviceInfoRepository struct {
	db *gorm.DB
}

func NewDeviceInfoRepository(
	db *gorm.DB,
) DeviceInfoRepository {
	return &deviceInfoRepository{
		db: db,
	}
}

var _ DeviceInfoRepository = (*deviceInfoRepository)(nil)

func (r *deviceInfoRepository) FirstOrCreateDeviceInfo(device *DeviceInfo) (*DeviceInfo, error) {
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
