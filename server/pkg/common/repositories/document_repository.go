package repositories

import "github.com/ahmola/Hanwoo-Tax-Corporation/pkg/common/models"

// Document
type DocumentRepository struct {
	GormRepository[models.Document]
}

func (r *DocumentRepository) FindByTitle(title string) ([]*models.Document, error) {
	var documents []*models.Document
	err := r.DB.Where("title LIKE ?", "%"+title+"%").Find(&documents).Error
	return documents, err
}

func (r *DocumentRepository) FindByCategory(category string) ([]*models.Document, error) {
	var documents []*models.Document
	err := r.DB.Where("category = ?", category).Find(&documents).Error
	return documents, err
}
