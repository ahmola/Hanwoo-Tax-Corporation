package document

type DocumentService struct {
	Repo *DocumentRepository
}

func (svc *DocumentService) CreateDocument(document *Document) error {
	return svc.Repo.Create(document)
}

func (svc *DocumentService) GetDocumentsByTitle(title string) ([]Document, error) {
	return svc.Repo.FindByTitle(title)
}

func (svc *DocumentService) GetDocumentsByCategory(category string) ([]Document, error) {
	return svc.Repo.FindByCategory(category)
}

func (svc *DocumentService) UpdateDocument(document *Document) error {
	return svc.Repo.Update(document)
}

func (svc *DocumentService) DeleteDocument(id uint) error {
	return svc.Repo.Delete(id)
}
