package services

import (
	errors "github.com/ahmola/Hanwoo-Tax-Corporation/pkg/common/errors"
	models "github.com/ahmola/Hanwoo-Tax-Corporation/pkg/common/models"
	repositories "github.com/ahmola/Hanwoo-Tax-Corporation/pkg/common/repositories"
)

type NoticeService struct {
	Repo *repositories.NoticeRepository
}

var serviceName = "NoticeService"

func (svc *NoticeService) CreateNotice(notice *models.Notice) error {
	var functionName = "CreateNotice"
	if err := svc.Repo.Create(notice); err != nil {
		return errors.CommitError(serviceName, functionName,
			"error occured during creating notice", err)
	}

	return nil
}

func (svc *NoticeService) GetNoticesByTitle(title string) ([]models.Notice, error) {
	var functionName = "GetNoticesByTitle"
	notices, err := svc.Repo.FindByTitle(title)
	if err != nil {
		return nil, errors.CommitError(serviceName, functionName,
			"error occured during finding notices by title", err)
	}
	return notices, nil
}

func (svc *NoticeService) GetNoticesByContent(content string) ([]models.Notice, error) {
	var functionName = "GetNoticesByContent"
	notices, err := svc.Repo.FindByContent(content)
	if err != nil {
		return nil, errors.CommitError(serviceName, functionName,
			"error occured during finding notices by content", err)
	}
	return notices, nil
}

func (svc *NoticeService) GetNoticesByCategory(category string) ([]models.Notice, error) {
	var functionName = "GetNoticesByCategory"
	notices, err := svc.Repo.FindByCategory(category)
	if err != nil {
		return nil, errors.CommitError(serviceName, functionName,
			"error occured during finding notices by category", err)
	}
	return notices, nil
}

func (svc *NoticeService) GetNoticesByIsImportant(isImportant bool) ([]models.Notice, error) {
	var functionName = "GetNoticesByIsImportant"
	notices, err := svc.Repo.FindByIsImportant(isImportant)

	if err != nil {
		return nil, errors.CommitError(serviceName, functionName,
			"error occured during finding notices by isImportant", err)
	}
	return notices, nil
}

func (svc *NoticeService) UpdateNotice(notice *models.Notice) error {
	var functionName = "UpdateNotice"
	if err := svc.Repo.Update(notice); err != nil {
		return errors.CommitError(serviceName, functionName,
			"error occured during updating notice", err)
	}
	return nil
}

func (svc *NoticeService) DeleteNotice(id uint) error {
	var functionName = "DeleteNotice"
	if err := svc.Repo.Delete(id); err != nil {
		return errors.CommitError(serviceName, functionName,
			"error occured during deleting notice", err)
	}
	return nil
}
