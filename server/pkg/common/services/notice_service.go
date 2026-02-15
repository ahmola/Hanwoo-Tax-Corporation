package services

import (
	errors "github.com/ahmola/Hanwoo-Tax-Corporation/pkg/common/errors"
	models "github.com/ahmola/Hanwoo-Tax-Corporation/pkg/common/models"
	repositories "github.com/ahmola/Hanwoo-Tax-Corporation/pkg/common/repositories"
	"github.com/ahmola/Hanwoo-Tax-Corporation/pkg/common/utils"
)

type NoticeService struct {
	Repo *repositories.NoticeRepository
}

var noticeServiceName = "NoticeService"

func (svc *NoticeService) CreateNotice(notice *models.Notice) (*models.Notice, error) {
	var functionName = "CreateNotice"

	if err := svc.Repo.Create(notice); err != nil {
		return nil, errors.CommitError(noticeServiceName, functionName,
			"error occured during creating notice", err)
	}

	return notice, nil
}

func (svc *NoticeService) GetAllNotices() ([]*models.Notice, error) {
	var functionName = "GetAllNotices"

	notices, err := svc.Repo.FindAll()
	if err != nil {
		return nil, errors.CommitError(noticeServiceName, functionName,
			"error occured during finding all notices", err)
	}

	return notices, nil
}

func (svc *NoticeService) GetNoticeByID(id string) (*models.Notice, error) {
	var functionName = "GetNoticeByID"

	noticeID, err := utils.ConvertStringToUint(id)
	if err != nil {
		return nil, errors.CommitError(noticeServiceName, functionName,
			"error occured during converting string to uint", err)
	}

	notice, err := svc.Repo.FindByID(noticeID)
	if err != nil {
		return nil, errors.CommitError(noticeServiceName, functionName,
			"error occured during finding notice by id", err)
	}

	return notice, nil
}

func (svc *NoticeService) GetNoticesByTitle(title string) ([]*models.Notice, error) {
	var functionName = "GetNoticesByTitle"

	notices, err := svc.Repo.FindByTitle(title)
	if err != nil {
		return nil, errors.CommitError(noticeServiceName, functionName,
			"error occured during finding notices by title", err)
	}

	return notices, nil
}

func (svc *NoticeService) GetNoticesByContent(content string) ([]*models.Notice, error) {
	var functionName = "GetNoticesByContent"

	notices, err := svc.Repo.FindByContent(content)
	if err != nil {
		return nil, errors.CommitError(noticeServiceName, functionName,
			"error occured during finding notices by content", err)
	}
	return notices, nil
}

func (svc *NoticeService) GetNoticesByCategory(category string) ([]*models.Notice, error) {
	var functionName = "GetNoticesByCategory"

	notices, err := svc.Repo.FindByCategory(category)
	if err != nil {
		return nil, errors.CommitError(noticeServiceName, functionName,
			"error occured during finding notices by category", err)
	}

	return notices, nil
}

func (svc *NoticeService) GetNoticesByIsImportant(isImportant string) ([]*models.Notice, error) {
	var functionName = "GetNoticesByIsImportant"

	important, err := utils.ConvertStringToBool(isImportant)

	notices, err := svc.Repo.FindByIsImportant(important)
	if err != nil {
		return nil, errors.CommitError(noticeServiceName, functionName,
			"error occured during finding notices by isImportant", err)
	}

	return notices, nil
}

func (svc *NoticeService) UpdateNotice(notice *models.Notice) (*models.Notice, error) {
	var functionName = "UpdateNotice"

	if err := svc.Repo.Update(notice); err != nil {
		return nil, errors.CommitError(noticeServiceName, functionName,
			"error occured during updating notice", err)
	}

	return notice, nil
}

func (svc *NoticeService) DeleteNotice(id string) error {
	var functionName = "DeleteNotice"

	noticeID, err := utils.ConvertStringToUint(id)
	if err != nil {
		return errors.CommitError(noticeServiceName, functionName,
			"error occured during converting string to uint", err)
	}

	if err := svc.Repo.Delete(noticeID); err != nil {
		return errors.CommitError(noticeServiceName, functionName,
			"error occured during deleting notice", err)
	}

	return nil
}
