package common

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
	return r.DB.Create(entity).Error
}

func (r *GormRepository[T]) FindByID(id uint) (*T, error) {
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
	return r.DB.Save(entity).Error
}

func (r *GormRepository[T]) Delete(id uint) error {
	var entity T
	return r.DB.Delete(&entity, id).Error
}

// Notice
type NoticeRepository struct {
	GormRepository[Notice]
}

func (r *NoticeRepository) FindByTitle(title string) ([]Notice, error) {
	var notices []Notice
	err := r.DB.Where("title LIKE ?", "%"+title+"%").Find(&notices).Error
	return notices, err
}

func (r *NoticeRepository) FindByContent(content string) ([]Notice, error) {
	var notices []Notice
	err := r.DB.Where("content LIKE ?", "%"+content+"%").Find(&notices).Error
	return notices, err
}

func (r *NoticeRepository) FindByCategory(category string) ([]Notice, error) {
	var notices []Notice
	err := r.DB.Where("category = ?", category).Find(&notices).Error
	return notices, err
}

func (r *NoticeRepository) FindByIsImportant(isImportant bool) ([]Notice, error) {
	var notices []Notice
	err := r.DB.Where("is_important = ?", isImportant).Find(&notices).Error
	return notices, err
}

// Document
type DocumentRepository struct {
	GormRepository[Document]
}

func (r *DocumentRepository) FindByTitle(title string) ([]Document, error) {
	var documents []Document
	err := r.DB.Where("title LIKE ?", "%"+title+"%").Find(&documents).Error
	return documents, err
}

func (r *DocumentRepository) FindByCategory(category string) ([]Document, error) {
	var documents []Document
	err := r.DB.Where("category = ?", category).Find(&documents).Error
	return documents, err
}

// Contact
type ContactRepository struct {
	GormRepository[Contact]
}
