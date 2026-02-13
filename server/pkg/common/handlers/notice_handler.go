package handlers

import (
	models "github.com/ahmola/Hanwoo-Tax-Corporation/pkg/common/models"
	services "github.com/ahmola/Hanwoo-Tax-Corporation/pkg/common/services"
)

type NoticeHandler struct {
	Svc *services.NoticeService
}

func (hdl *NoticeHandler) CreateNotice(notice *models.Notice) error {
	return hdl.Svc.CreateNotice(notice)
}

func (hdl *NoticeHandler) GetNoticesByTitle(title string) ([]models.Notice, error) {
	return hdl.Svc.GetNoticesByTitle(title)
}

func (hdl *NoticeHandler) GetNoticesByContent(content string) ([]models.Notice, error) {
	return hdl.Svc.GetNoticesByContent(content)
}

func (hdl *NoticeHandler) GetNoticesByCategory(category string) ([]models.Notice, error) {
	return hdl.Svc.GetNoticesByCategory(category)
}

func (hdl *NoticeHandler) GetNoticesByIsImportant(isImportant bool) ([]models.Notice, error) {
	return hdl.Svc.GetNoticesByIsImportant(isImportant)
}

func (hdl *NoticeHandler) UpdateNotice(notice *models.Notice) error {
	return hdl.Svc.UpdateNotice(notice)
}

func (hdl *NoticeHandler) DeleteNotice(id uint) error {
	return hdl.Svc.DeleteNotice(id)
}
