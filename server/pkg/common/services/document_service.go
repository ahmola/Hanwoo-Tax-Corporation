package services

import (
	errors "github.com/ahmola/Hanwoo-Tax-Corporation/pkg/common/errors"
	models "github.com/ahmola/Hanwoo-Tax-Corporation/pkg/common/models"
	repositories "github.com/ahmola/Hanwoo-Tax-Corporation/pkg/common/repositories"
	"github.com/ahmola/Hanwoo-Tax-Corporation/pkg/common/utils"
)

type DocumentService struct {
	Repo *repositories.DocumentRepository
}

var documentServiceName = "DocumentService"

func (svc *DocumentService) CreateDocument(document *models.Document) (*models.Document, error) {
	var functionName = "CreateDocument"

	if err := svc.Repo.Create(document); err != nil {
		return nil, errors.CommitError(documentServiceName, functionName,
			"error occured during creating document", err)
	}

	return document, nil
}

func (svc *DocumentService) GetDocumentByID(id string) (*models.Document, error) {
	var functionName = "GetDocumentByID"

	documentID, err := utils.ConvertStringToUint(id)
	if err != nil {
		return nil, errors.CommitError(documentServiceName, functionName,
			"error occured during converting string to uint", err)
	}

	document, err := svc.Repo.FindByID(documentID)
	if err != nil {
		return nil, errors.CommitError(documentServiceName, functionName,
			"error occured during finding notice by id", err)
	}

	return document, nil
}

func (svc *DocumentService) GetAllDocuments() ([]*models.Document, error) {
	var functionName = "GetAllDocuments"

	documents, err := svc.Repo.FindAll()
	if err != nil {
		return nil, errors.CommitError(documentServiceName, functionName,
			"error occured during finding documents", err)
	}

	return documents, nil
}

func (svc *DocumentService) GetDocumentsByTitle(title string) ([]*models.Document, error) {
	var functionName = "GetDocumentsByTitle"

	documents, err := svc.Repo.FindByTitle(title)
	if err != nil {
		return nil, errors.CommitError(documentServiceName, functionName,
			"error occured during reading documents", err)
	}

	return documents, nil
}

func (svc *DocumentService) GetDocumentsByCategory(category string) ([]*models.Document, error) {
	var functionName = "GetDocumentsByCategory"

	documents, err := svc.Repo.FindByCategory(category)
	if err != nil {
		return nil, errors.CommitError(documentServiceName, functionName,
			"error occured during reading documents", err)
	}

	return documents, err
}

func (svc *DocumentService) UpdateDocument(document *models.Document) (*models.Document, error) {
	var functionName = "UpdatedDocument"

	if err := svc.Repo.Update(document); err != nil {
		return nil, errors.CommitError(documentServiceName, functionName,
			"error occured during updating document", err)
	}

	return document, nil
}

func (svc *DocumentService) DeleteDocument(id string) error {
	var functionName = "DeleteDocument"

	contactID, err := utils.ConvertStringToUint(id)
	if err != nil {
		return errors.CommitError(documentServiceName, functionName,
			"error occured during converting string to uint64", err)
	}

	if err := svc.Repo.Delete(contactID); err != nil {
		return errors.CommitError(documentServiceName, functionName,
			"error occured during deleting document", err)
	}

	return nil
}
