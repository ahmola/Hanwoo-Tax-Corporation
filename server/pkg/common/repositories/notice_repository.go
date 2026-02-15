package repositories

import "github.com/ahmola/Hanwoo-Tax-Corporation/pkg/common/models"

// Notice
type NoticeRepository struct {
	GormRepository[models.Notice]
}

func (r *NoticeRepository) FindByTitle(title string) ([]*models.Notice, error) {
	var notices []*models.Notice
	err := r.DB.Where("title LIKE ?", "%"+title+"%").Find(&notices).Error
	return notices, err
}

func (r *NoticeRepository) FindByContent(content string) ([]*models.Notice, error) {
	var notices []*models.Notice
	err := r.DB.Where("content LIKE ?", "%"+content+"%").Find(&notices).Error
	return notices, err
}

func (r *NoticeRepository) FindByCategory(category string) ([]*models.Notice, error) {
	var notices []*models.Notice
	err := r.DB.Where("category = ?", category).Find(&notices).Error
	return notices, err
}

func (r *NoticeRepository) FindByIsImportant(isImportant bool) ([]*models.Notice, error) {
	var notices []*models.Notice
	err := r.DB.Where("is_important = ?", isImportant).Find(&notices).Error
	return notices, err
}
