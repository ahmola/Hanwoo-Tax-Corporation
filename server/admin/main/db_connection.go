package main

import (
	"log/slog"
	"os"

	"github.com/ahmola/Hanwoo-Tax-Corporation/pkg/common"
	contact "github.com/ahmola/Hanwoo-Tax-Corporation/pkg/common/contact"
	document "github.com/ahmola/Hanwoo-Tax-Corporation/pkg/common/document"
	notice "github.com/ahmola/Hanwoo-Tax-Corporation/pkg/common/notice"
)

func dbInit() (*notice.NoticeHandler, *document.DocumentHandler, *contact.ContactHandler) {
	// DB 설정
	noticeDB := common.DBConfig("notice")
	if err := noticeDB.AutoMigrate(&notice.Notice{}); err != nil {
		slog.Error("Notice DB Migration Failed!", "error", err)
		os.Exit(1)
	}

	documentDB := common.DBConfig("document")
	if err := documentDB.AutoMigrate(&document.Document{}); err != nil {
		slog.Error("Document DB Migration Failed!", "error", err)
		os.Exit(1)
	}

	contactDB := common.DBConfig("contact")
	if err := contactDB.AutoMigrate(&contact.Contact{}); err != nil {
		slog.Error("Contact DB Migration Failed!", "error", err)
		os.Exit(1)
	}

	// 의존성 주입
	// Notice
	slog.Info("Injecting Notice Repository Dependency")
	noticeRepo := &notice.NoticeRepository{GormRepository: common.GormRepository[notice.Notice]{DB: noticeDB}}
	slog.Info("Notice Repository is Ready : ", "db_name", noticeRepo.DB.Name())

	noticeSvc := &notice.NoticeService{Repo: noticeRepo}
	slog.Info("Notice Service is Ready : ", "db_name", noticeSvc.Repo.DB.Name())

	noticeHdl := &notice.NoticeHandler{Svc: noticeSvc}
	slog.Info("Notice Handler is Ready : ", "db_name", noticeHdl.Svc.Repo.DB.Name())

	// Document
	slog.Info("Injecting Document Repository Dependency")
	documentRepo := &document.DocumentRepository{GormRepository: common.GormRepository[document.Document]{DB: documentDB}}
	slog.Info("Document Repository is Ready : ", "db_name", documentRepo.DB.Name())

	documentSvc := &document.DocumentService{Repo: documentRepo}
	slog.Info("Document Service is Ready : ", "db_name", documentSvc.Repo.DB.Name())

	documentHdl := &document.DocumentHandler{Svc: documentSvc}
	slog.Info("Document Handler is Ready : ", "db_name", documentHdl.Svc.Repo.DB.Name())

	// Contact
	slog.Info("Injecting Contact Repository Dependency")
	contactRepo := &contact.ContactRepository{GormRepository: common.GormRepository[contact.Contact]{DB: contactDB}}
	slog.Info("Contact Repository is Ready : ", "db_name", contactRepo.DB.Name())

	contactSvc := &contact.ContactService{Repo: contactRepo}
	slog.Info("Contact Service is Ready : ", "db_name", contactSvc.Repo.DB.Name())

	contactHdl := &contact.ContactHandler{Svc: contactSvc}
	slog.Info("Contact Handler is Ready : ", "db_name", contactHdl.Svc.Repo.DB.Name())

	return noticeHdl, documentHdl, contactHdl
}
