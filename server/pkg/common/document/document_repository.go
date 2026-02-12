package document

import "github.com/ahmola/Hanwoo-Tax-Corporation/pkg/common"

// Document
type DocumentRepository struct {
	common.GormRepository[Document]
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
