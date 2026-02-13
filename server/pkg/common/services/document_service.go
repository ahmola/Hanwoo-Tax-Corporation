package services

import (
	models "github.com/ahmola/Hanwoo-Tax-Corporation/pkg/common/models"
	repositories "github.com/ahmola/Hanwoo-Tax-Corporation/pkg/common/repositories"
)

type DocumentService struct {
	Repo *repositories.DocumentRepository
}

func (svc *DocumentService) CreateDocument(document *models.Document) error {
	return svc.Repo.Create(document)
}

func (svc *DocumentService) GetDocumentsByTitle(title string) ([]models.Document, error) {
	return svc.Repo.FindByTitle(title)
}

func (svc *DocumentService) GetDocumentsByCategory(category string) ([]models.Document, error) {
	return svc.Repo.FindByCategory(category)
}

func (svc *DocumentService) UpdateDocument(document *models.Document) error {
	return svc.Repo.Update(document)
}

func (svc *DocumentService) DeleteDocument(id uint) error {
	return svc.Repo.Delete(id)
}
