package contact

import "github.com/ahmola/Hanwoo-Tax-Corporation/pkg/common"

// Contact
type ContactRepository struct {
	common.GormRepository[Contact]
}

func (r *ContactRepository) FindByName(name string) ([]Contact, error) {
	var contacts []Contact
	err := r.DB.Where("name LIKE ?", "%"+name+"%").Find(&contacts).Error
	return contacts, err
}

func (r *ContactRepository) FindByPhone(phone string) ([]Contact, error) {
	var contacts []Contact
	err := r.DB.Where("phone LIKE ?", "%"+phone+"%").Find(&contacts).Error
	return contacts, err
}

func (r *ContactRepository) FindByMessage(message string) ([]Contact, error) {
	var contacts []Contact
	err := r.DB.Where("message LIKE ?", "%"+message+"%").Find(&contacts).Error
	return contacts, err
}
