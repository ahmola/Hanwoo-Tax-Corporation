package handlers

import (
	"log/slog"

	errors "github.com/ahmola/Hanwoo-Tax-Corporation/pkg/common/errors"
	models "github.com/ahmola/Hanwoo-Tax-Corporation/pkg/common/models"
	services "github.com/ahmola/Hanwoo-Tax-Corporation/pkg/common/services"
	"github.com/gin-gonic/gin"
)

type ContactHandler struct {
	Svc *services.ContactService
}

var contactHandlerName = "ContactHandler"

func (hdl *ContactHandler) CreateContact(c *gin.Context) {
	var functionName = "CreateContact"
	var req models.Contact

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": errors.CommitError(contactHandlerName, functionName,
			"error occured during extracting request", err)})
		return
	}

	res, err := hdl.Svc.CreateContact(&req)
	if err != nil {
		c.JSON(500, gin.H{"error": errors.CommitError(contactHandlerName, functionName,
			"error occured during creating contact", err)})
		return
	}

	c.JSON(201, res)
}

func (hdl *ContactHandler) GetAllContacts(c *gin.Context) {
	var functionName = "GetAllContacts"

	res, err := hdl.Svc.GetAllContacts()
	if err != nil {
		c.JSON(500, gin.H{"error": errors.CommitError(documentHandlerName, functionName,
			"error occured during reading documents", err)})
		return
	}

	c.JSON(200, res)
}

func (hdl *ContactHandler) GetCotactsByID(c *gin.Context) {
	var functionName = "GetContactsByID"
	id := c.Param("id")
	slog.Info("ContactId: ", id)

	contact, err := hdl.Svc.GetContactsByID(id)
	if err != nil {
		c.JSON(500, gin.H{"error": errors.CommitError(contactHandlerName, functionName,
			"error occured during reading contacts", err)})
		return
	}

	c.JSON(200, contact)
}

func (hdl *ContactHandler) UpdateContact(c *gin.Context) {
	var functionName = "UpdateContact"
	var req models.Contact

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": errors.CommitError(contactHandlerName, functionName,
			"error occured during extracting request", err)})
		return
	}

	res, err := hdl.Svc.UpdateContact(&req)
	if err != nil {
		c.JSON(500, gin.H{"error": errors.CommitError(contactHandlerName, functionName,
			"error occured during updating contacts", err)})
		return
	}

	c.JSON(200, res)
}

func (hdl *ContactHandler) DeleteContact(c *gin.Context) {
	var functionName = "deletieContact"
	id := c.Param("id")

	err := hdl.Svc.DeleteContact(id)
	if err != nil {
		c.JSON(500, gin.H{"error": errors.CommitError(contactHandlerName, functionName,
			"error occured during deleting contact", err)})
		return
	}

	c.JSON(200, true)
}
