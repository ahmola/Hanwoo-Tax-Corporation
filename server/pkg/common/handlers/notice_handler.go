package handlers

import (
	errors "github.com/ahmola/Hanwoo-Tax-Corporation/pkg/common/errors"
	models "github.com/ahmola/Hanwoo-Tax-Corporation/pkg/common/models"
	services "github.com/ahmola/Hanwoo-Tax-Corporation/pkg/common/services"
	"github.com/gin-gonic/gin"
)

type NoticeHandler struct {
	Svc *services.NoticeService
}

var noticeHandlerName = "NoticeHandler"

func (hdl *NoticeHandler) CreateNotice(c *gin.Context) {
	var req models.Notice
	var functionName = "CreateNotice"

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": errors.CommitError(noticeHandlerName, functionName,
			"error occured during extracting request", err)})
		return
	}

	res, err := hdl.Svc.CreateNotice(&req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, res)
}

func (hdl *NoticeHandler) GetAllNotices(c *gin.Context) {
	var functionName = "GetAllNotices"

	res, err := hdl.Svc.GetAllNotices()
	if err != nil {
		c.JSON(500, gin.H{"error": errors.CommitError(noticeHandlerName, functionName,
			"error occured during reading notices", err)})
		return
	}

	c.JSON(200, res)
}

func (hdl *NoticeHandler) GetNoticeByID(c *gin.Context) {
	var functionName = "GetNoticeByID"
	id := c.Param("id")

	res, err := hdl.Svc.GetNoticeByID(id)
	if err != nil {
		c.JSON(500, gin.H{"error": errors.CommitError(noticeHandlerName, functionName,
			"error occured during reading notice", err)})
		return
	}

	c.JSON(200, res)
}

func (hdl *NoticeHandler) UpdateNotice(c *gin.Context) {
	var req models.Notice
	var functionName = "UpdateNotice"

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": errors.CommitError(noticeHandlerName, functionName,
			"error occured during extracting request", err)})
		return
	}

	res, err := hdl.Svc.UpdateNotice(&req)
	if err != nil {
		c.JSON(500, gin.H{"error": errors.CommitError(noticeHandlerName, functionName,
			"error occured during updating notice", err)})
		return
	}

	c.JSON(200, res)

}

func (hdl *NoticeHandler) DeleteNotice(c *gin.Context) {
	var functionName = "deleteNotice"
	id := c.Param("id")

	if err := hdl.Svc.DeleteNotice(id); err != nil {
		c.JSON(500, errors.CommitError(noticeHandlerName, functionName,
			"error occured during deleting notice", err))
		return
	}

	c.JSON(200, true)
}
