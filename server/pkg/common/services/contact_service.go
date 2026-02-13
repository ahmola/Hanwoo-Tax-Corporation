package services

import (
	models "github.com/ahmola/Hanwoo-Tax-Corporation/pkg/common/models"
	repositories "github.com/ahmola/Hanwoo-Tax-Corporation/pkg/common/repositories"
)

type ContactService struct {
	Repo *repositories.ContactRepository
}

func (svc *ContactService) CreateContact(contact *models.Contact) error {
	return svc.Repo.Create(contact)
}

func (svc *ContactService) GetContectsByID(id uint) (*models.Contact, error) {
	return svc.Repo.FindByID(id)
}

func (svc *ContactService) GetContactsByName(name string) ([]models.Contact, error) {
	return svc.Repo.FindByName(name)
}

func (svc *ContactService) GetContactsByPhone(phone string) ([]models.Contact, error) {
	return svc.Repo.FindByPhone(phone)
}

func (svc *ContactService) GetContactsByMessage(message string) ([]models.Contact, error) {
	return svc.Repo.FindByMessage(message)
}

func (svc *ContactService) UpdateContact(contact *models.Contact) error {
	return svc.Repo.Update(contact)
}

func (svc *ContactService) DeleteContact(id uint) error {
	return svc.Repo.Delete(id)
}
