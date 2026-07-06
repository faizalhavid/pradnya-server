package shared

import "gorm.io/gorm"

type BaseRepository[T any] interface {
	Create(entity *T) error
	FindByID(id string) (*T, error)
	FindByField(fieldName string, value interface{}) (*T, error)
	Update(entity *T) error
	SoftDelete(entity *T) error
	PermanentlyDelete(entity *T) error
}

type baseRepository[T any] struct {
	db *gorm.DB
}

func NewBaseRepository[T any](db *gorm.DB) BaseRepository[T] {
	return &baseRepository[T]{db: db}
}

func (r *baseRepository[T]) Create(entity *T) error {
	return r.db.Create(entity).Error
}

func (r *baseRepository[T]) FindByID(id string) (*T, error) {
	var entity T
	err := r.db.Where("id=?", id).First(&entity).Error
	return &entity, err
}

func (r *baseRepository[T]) FindByField(fieldName string, value interface{}) (*T, error) {
	var entity T
	err := r.db.Where(fieldName+"=?", value).First(&entity).Error
	return &entity, err
}

func (r *baseRepository[T]) Update(entity *T) error {
	return r.db.Save(entity).Error
}

func (r *baseRepository[T]) SoftDelete(entity *T) error {
	return r.db.Where("id=?", entity).Update("is_deleted", true).Error
}

func (r *baseRepository[T]) PermanentlyDelete(entity *T) error {
	return r.db.Delete(entity).Error
}
