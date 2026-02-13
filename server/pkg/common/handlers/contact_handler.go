package handlers

import (
	models "github.com/ahmola/Hanwoo-Tax-Corporation/pkg/common/models"
	services "github.com/ahmola/Hanwoo-Tax-Corporation/pkg/common/services"
)

type ContactHandler struct {
	Svc *services.ContactService
}

func (hdl *ContactHandler) CreateContact(contact *models.Contact) error {
	return hdl.Svc.CreateContact(contact)
}

func (hdl *ContactHandler) GetCotectsByID(id uint) (*models.Contact, error) {
	return hdl.Svc.GetContectsByID(id)
}

func (hdl *ContactHandler) GetContactsByName(name string) ([]models.Contact, error) {
	return hdl.Svc.GetContactsByName(name)
}

func (hdl *ContactHandler) GetContactsByPhone(phone string) ([]models.Contact, error) {
	return hdl.Svc.GetContactsByPhone(phone)
}

func (hdl *ContactHandler) GetContactsByMessage(message string) ([]models.Contact, error) {
	return hdl.Svc.GetContactsByMessage(message)
}

func (hdl *ContactHandler) UpdateContact(contact *models.Contact) error {
	return hdl.Svc.UpdateContact(contact)
}

func (hdl *ContactHandler) DeleteContact(id uint) error {
	return hdl.Svc.DeleteContact(id)
}
