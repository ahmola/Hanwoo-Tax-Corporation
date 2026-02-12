package notice

type NoticeHandler struct {
	Svc *NoticeService
}

func (hdl *NoticeHandler) CreateNotice(notice *Notice) error {
	return hdl.Svc.CreateNotice(notice)
}

func (hdl *NoticeHandler) GetNoticesByTitle(title string) ([]Notice, error) {
	return hdl.Svc.GetNoticesByTitle(title)
}

func (hdl *NoticeHandler) GetNoticesByContent(content string) ([]Notice, error) {
	return hdl.Svc.GetNoticesByContent(content)
}

func (hdl *NoticeHandler) GetNoticesByCategory(category string) ([]Notice, error) {
	return hdl.Svc.GetNoticesByCategory(category)
}

func (hdl *NoticeHandler) GetNoticesByIsImportant(isImportant bool) ([]Notice, error) {
	return hdl.Svc.GetNoticesByIsImportant(isImportant)
}

func (hdl *NoticeHandler) UpdateNotice(notice *Notice) error {
	return hdl.Svc.UpdateNotice(notice)
}

func (hdl *NoticeHandler) DeleteNotice(id uint) error {
	return hdl.Svc.DeleteNotice(id)
}
