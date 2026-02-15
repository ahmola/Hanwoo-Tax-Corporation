package repositories

import "github.com/ahmola/Hanwoo-Tax-Corporation/pkg/common/models"

// Contact
type ContactRepository struct {
	GormRepository[models.Contact]
}

func (r *ContactRepository) FindByName(name string) ([]*models.Contact, error) {
	var contacts []*models.Contact
	err := r.DB.Where("name LIKE ?", "%"+name+"%").Find(&contacts).Error
	return contacts, err
}

func (r *ContactRepository) FindByPhone(phone string) ([]*models.Contact, error) {
	var contacts []*models.Contact
	err := r.DB.Where("phone LIKE ?", "%"+phone+"%").Find(&contacts).Error
	return contacts, err
}

func (r *ContactRepository) FindByMessage(message string) ([]*models.Contact, error) {
	var contacts []*models.Contact
	err := r.DB.Where("message LIKE ?", "%"+message+"%").Find(&contacts).Error
	return contacts, err
}
