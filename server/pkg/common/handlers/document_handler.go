package handlers

import (
	models "github.com/ahmola/Hanwoo-Tax-Corporation/pkg/common/models"
	services "github.com/ahmola/Hanwoo-Tax-Corporation/pkg/common/services"

	"github.com/gin-gonic/gin"
)

type DocumentHandler struct {
	Svc *services.DocumentService
}

func (hdl *DocumentHandler) CreateDocument(c *gin.Context) {
	var req models.Document
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}

	res, err := hdl.Svc.CreateDocument(req)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(201, res)
}

func (hdl *DocumentHandler) GetDocumentsByTitle(title string) ([]models.Document, error) {
	return hdl.Svc.GetDocumentsByTitle(title)
}

func (hdl *DocumentHandler) GetDocumentsByCategory(category string) ([]models.Document, error) {
	return hdl.Svc.GetDocumentsByCategory(category)
}

func (hdl *DocumentHandler) UpdateDocument(document *models.Document) error {
	return hdl.Svc.UpdateDocument(document)
}

func (hdl *DocumentHandler) DeleteDocument(id uint) error {
	return hdl.Svc.DeleteDocument(id)
}
