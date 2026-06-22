package shared

import (
	"gorm.io/gorm"
)

type BaseModel struct {
	ID string `gorm:"type:char(26);primaryKey" json:"id"`
	gorm.Model
	IsDeleted bool            `gorm:"default:false" json:"is_deleted"`
	DeletedAt *gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (e *BaseModel) BeforeCreate(tx *gorm.DB) error {
	if e.ID == "" {
		id, err := New() // ulid
		if err != nil {
			return err
		}
		e.ID = id
	}
	return nil
}
