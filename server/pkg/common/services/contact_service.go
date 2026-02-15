package services

import (
	errors "github.com/ahmola/Hanwoo-Tax-Corporation/pkg/common/errors"
	models "github.com/ahmola/Hanwoo-Tax-Corporation/pkg/common/models"
	repositories "github.com/ahmola/Hanwoo-Tax-Corporation/pkg/common/repositories"
	"github.com/ahmola/Hanwoo-Tax-Corporation/pkg/common/utils"
)

type ContactService struct {
	Repo *repositories.ContactRepository
}

var contactServiceName = "ContactService"

func (svc *ContactService) CreateContact(contact *models.Contact) (*models.Contact, error) {
	var functionName = "CreateContact"

	if err := svc.Repo.Create(contact); err != nil {
		return nil, errors.CommitError(contactServiceName, functionName,
			"error occured during creating contact", err)
	}

	return contact, nil
}

func (svc *ContactService) GetAllContacts() ([]*models.Contact, error) {
	var functionName = "GetAllContacts"

	contacts, err := svc.Repo.FindAll()
	if err != nil {
		return nil, errors.CommitError(contactServiceName, functionName,
			"error occured during reading contacts", err)
	}

	return contacts, nil
}

func (svc *ContactService) GetContactsByID(id string) (*models.Contact, error) {
	var functionName = "GetContactsByID"

	contactID, err := utils.ConvertStringToUint(id)
	if err != nil {
		return nil, errors.CommitError(contactServiceName, functionName,
			"error occured during converting string to uint", err)
	}

	contact, err := svc.Repo.FindByID(contactID)
	if err != nil {
		return nil, errors.CommitError(contactServiceName, functionName,
			"error occured during getting contact by id", err)
	}

	return contact, nil
}

func (svc *ContactService) GetContactsByName(name string) ([]*models.Contact, error) {
	var functionName = "GetContactsByName"

	contacts, err := svc.Repo.FindByName(name)
	if err != nil {
		return nil, errors.CommitError(contactServiceName, functionName,
			"error occured during getting contact by name", err)
	}

	return contacts, nil
}

func (svc *ContactService) GetContactsByPhone(phone string) ([]*models.Contact, error) {
	var functionName = "GetContactsByPhone"

	contacs, err := svc.Repo.FindByPhone(phone)
	if err != nil {
		return nil, errors.CommitError(contactServiceName, functionName,
			"error occured during getting contact by phone", err)
	}

	return contacs, nil
}

func (svc *ContactService) GetContactsByMessage(message string) ([]*models.Contact, error) {
	var functionName = "GetContactsByMessage"

	contacts, err := svc.Repo.FindByMessage(message)
	if err != nil {
		return nil, errors.CommitError(contactServiceName, functionName,
			"error occured during reading contact by message", err)
	}

	return contacts, nil
}

func (svc *ContactService) UpdateContact(contact *models.Contact) (*models.Contact, error) {
	var functionName = "UpdateContact"

	if err := svc.Repo.Update(contact); err != nil {
		return nil, errors.CommitError(contactServiceName, functionName,
			"error occured during updating contact", err)
	}

	return contact, nil
}

func (svc *ContactService) DeleteContact(id string) error {
	var functionName = "DeleteContact"

	contactID, err := utils.ConvertStringToUint(id)
	if err != nil {
		return errors.CommitError(noticeServiceName, functionName,
			"error occured during converting string to uint", err)
	}

	if err := svc.Repo.Delete(contactID); err != nil {
		return errors.CommitError(noticeServiceName, functionName,
			"error occured during deleting contact", err)
	}

	return nil
}
