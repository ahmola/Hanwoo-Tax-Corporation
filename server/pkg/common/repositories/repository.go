package repositories

import (
	"gorm.io/gorm"
)

// Repository Interfcae
type Repository[T any] interface {
	// create
	Create(entity *T) error

	// read
	FindByID(id uint) (*T, error)
	FindAll([]*T, error)

	// update
	Update(entity *T) error

	// delete
	Delete(id uint) error
}

// Gin Base ORM Repository
type GormRepository[T any] struct {
	Repository[T]
	DB *gorm.DB
}

func (r *GormRepository[T]) Create(entity *T) error {
	return r.DB.Create(&entity).Error
}

func (r *GormRepository[T]) FindByID(id uint64) (*T, error) {
	var entity T
	err := r.DB.First(&entity, id).Error
	return &entity, err
}

func (r *GormRepository[T]) FindAll() ([]*T, error) {
	var entities []*T
	err := r.DB.Find(&entities).Error
	return entities, err
}

func (r *GormRepository[T]) Update(entity *T) error {
	return r.DB.Save(&entity).Error
}

func (r *GormRepository[T]) Delete(id uint64) error {
	var entity T
	return r.DB.Delete(&entity, id).Error
}
