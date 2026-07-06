package shared

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        string `gorm:"type:char(26);primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
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
