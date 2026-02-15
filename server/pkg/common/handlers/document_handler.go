package handlers

import (
	errors "github.com/ahmola/Hanwoo-Tax-Corporation/pkg/common/errors"
	models "github.com/ahmola/Hanwoo-Tax-Corporation/pkg/common/models"
	services "github.com/ahmola/Hanwoo-Tax-Corporation/pkg/common/services"

	"github.com/gin-gonic/gin"
)

type DocumentHandler struct {
	Svc *services.DocumentService
}

var documentHandlerName = "DocumentHandler"

func (hdl *DocumentHandler) CreateDocument(c *gin.Context) {
	var functionName = "CreateDocument"
	var req models.Document

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": errors.CommitError(documentHandlerName, functionName,
			"error occured during extracting request", err)})
		return
	}

	res, err := hdl.Svc.CreateDocument(&req)
	if err != nil {
		c.JSON(500, gin.H{"error": errors.CommitError(documentHandlerName, functionName,
			"error occured during creaeting document", err)})
		return
	}

	c.JSON(201, res)
}

func (hdl *DocumentHandler) GetAllDocuments(c *gin.Context) {
	var functionName = "GetAllDocuments"

	res, err := hdl.Svc.GetAllDocuments()
	if err != nil {
		c.JSON(500, gin.H{"error": errors.CommitError(documentHandlerName, functionName,
			"error occured during reading documents", err)})
		return
	}

	c.JSON(200, res)
}

func (hdl *DocumentHandler) GetDocumentsByID(c *gin.Context) {
	var functionName = "GetDocumentsByID"
	id := c.Param("id")

	res, err := hdl.Svc.GetDocumentByID(id)
	if err != nil {
		c.JSON(500, gin.H{"error": errors.CommitError(documentHandlerName, functionName,
			"error occured during reading documents", err)})
		return
	}

	c.JSON(200, res)
}

func (hdl *DocumentHandler) UpdateDocument(c *gin.Context) {
	var functionName = "UpdatingDocument"
	var req models.Document

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": errors.CommitError(documentHandlerName, functionName,
			"error occured during extracting request", err)})
		return
	}

	res, err := hdl.Svc.UpdateDocument(&req)
	if err != nil {
		c.JSON(500, gin.H{"error": errors.CommitError(documentHandlerName, functionName,
			"error occured during updating documents", err)})
		return
	}

	c.JSON(200, res)
}

func (hdl *DocumentHandler) DeleteDocument(c *gin.Context) {
	var functionName = "DeleteDocument"
	id := c.Param("id")

	err := hdl.Svc.DeleteDocument(id)
	if err != nil {
		c.JSON(500, gin.H{"error": errors.CommitError(documentHandlerName, functionName,
			"error occured during deleting documents", err)})
		return
	}

	c.JSON(200, true)
}
