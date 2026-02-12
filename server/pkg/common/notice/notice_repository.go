package notice

import "github.com/ahmola/Hanwoo-Tax-Corporation/pkg/common"

// Notice
type NoticeRepository struct {
	common.GormRepository[Notice]
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
