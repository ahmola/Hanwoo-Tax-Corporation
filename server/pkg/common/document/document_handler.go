package document

type DocumentHandler struct {
	Svc *DocumentService
}

func (hdl *DocumentHandler) CreateDocument(document *Document) error {
	return hdl.Svc.CreateDocument(document)
}

func (hdl *DocumentHandler) GetDocumentsByTitle(title string) ([]Document, error) {
	return hdl.Svc.GetDocumentsByTitle(title)
}

func (hdl *DocumentHandler) GetDocumentsByCategory(category string) ([]Document, error) {
	return hdl.Svc.GetDocumentsByCategory(category)
}

func (hdl *DocumentHandler) UpdateDocument(document *Document) error {
	return hdl.Svc.UpdateDocument(document)
}

func (hdl *DocumentHandler) DeleteDocument(id uint) error {
	return hdl.Svc.DeleteDocument(id)
}
