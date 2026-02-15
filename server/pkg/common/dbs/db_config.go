package db

import (
	"log/slog"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	handlers "github.com/ahmola/Hanwoo-Tax-Corporation/pkg/common/handlers"
	models "github.com/ahmola/Hanwoo-Tax-Corporation/pkg/common/models"
	repositories "github.com/ahmola/Hanwoo-Tax-Corporation/pkg/common/repositories"
	services "github.com/ahmola/Hanwoo-Tax-Corporation/pkg/common/services"
)

var DB *gorm.DB

func DBConfig(serviceName string) *gorm.DB {
	// MySQL 주소
	// Compose의 환경변수에서 가져옴
	destination := os.Getenv("DB_DSN")
	if destination == "" {
		destination = "root:1234@tcp(127.0.0.1:3306)/" + serviceName + "?useUnicode=true&characterEncoding=utf8mb4&serverTimezone=UTC"
	}
	slog.Info("DB Destination : ", destination)

	// 데이터베이스 연결
	var err error

	slog.Info("Start Connection")
	for i := 0; i < 10; i++ {
		slog.Info("Try Connection... #", i+1)
		DB, err = gorm.Open(mysql.Open(destination), &gorm.Config{})

		if err == nil {
			break
		}

		slog.Warn("DB Connection failed, Restart... ", "error", err)
		time.Sleep(time.Second)
	}

	if err != nil || DB == nil {
		slog.Error("Eventually DB Connection failed.")
		os.Exit(1)
	}
	slog.Info(serviceName + " DB Connected! Start migrate")

	return DB
}

// DB 설정 & 의존성 주입
func DBInit() (*handlers.NoticeHandler, *handlers.DocumentHandler, *handlers.ContactHandler) {
	// Notice
	noticeDB := DBConfig("notice")
	if err := noticeDB.AutoMigrate(&models.Notice{}); err != nil {
		slog.Error("Notice DB Migration Failed!", "error", err)
		os.Exit(1)
	}

	slog.Info("Injecting Notice Repository Dependency")
	noticeRepo := &repositories.NoticeRepository{GormRepository: repositories.GormRepository[models.Notice]{DB: noticeDB}}
	slog.Info("Notice Repository is Ready : ", "db_name", noticeRepo.DB.Name())

	noticeSvc := &services.NoticeService{Repo: noticeRepo}
	slog.Info("Notice Service is Ready : ", "db_name", noticeSvc.Repo.DB.Name())

	noticeHdl := &handlers.NoticeHandler{Svc: noticeSvc}
	slog.Info("Notice Handler is Ready : ", "db_name", noticeHdl.Svc.Repo.DB.Name())

	// Document
	documentDB := DBConfig("document")
	if err := documentDB.AutoMigrate(&models.Document{}); err != nil {
		slog.Error("Document DB Migration Failed!", "error", err)
		os.Exit(1)
	}

	slog.Info("Injecting Document Repository Dependency")
	documentRepo := &repositories.DocumentRepository{GormRepository: repositories.GormRepository[models.Document]{DB: documentDB}}
	slog.Info("Document Repository is Ready : ", "db_name", documentRepo.DB.Name())

	documentSvc := &services.DocumentService{Repo: documentRepo}
	slog.Info("Document Service is Ready : ", "db_name", documentSvc.Repo.DB.Name())

	documentHdl := &handlers.DocumentHandler{Svc: documentSvc}
	slog.Info("Document Handler is Ready : ", "db_name", documentHdl.Svc.Repo.DB.Name())

	// Contact
	contactDB := DBConfig("contact")
	if err := contactDB.AutoMigrate(&models.Contact{}); err != nil {
		slog.Error("Contact DB Migration Failed!", "error", err)
		os.Exit(1)
	}

	slog.Info("Injecting Contact Repository Dependency")
	contactRepo := &repositories.ContactRepository{GormRepository: repositories.GormRepository[models.Contact]{DB: contactDB}}
	slog.Info("Contact Repository is Ready : ", "db_name", contactRepo.DB.Name())

	contactSvc := &services.ContactService{Repo: contactRepo}
	slog.Info("Contact Service is Ready : ", "db_name", contactSvc.Repo.DB.Name())

	contactHdl := &handlers.ContactHandler{Svc: contactSvc}
	slog.Info("Contact Handler is Ready : ", "db_name", contactHdl.Svc.Repo.DB.Name())

	return noticeHdl, documentHdl, contactHdl
}
