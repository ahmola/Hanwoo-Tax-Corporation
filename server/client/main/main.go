package main

import (
	"os"

	"log/slog"

	"github.com/gin-gonic/gin"

	dbs "github.com/ahmola/Hanwoo-Tax-Corporation/pkg/common/dbs"
)

// @title				Client API
// @version				1.0
// @description			API built with Gin
// @host				localhost:8082
// @BasePath			/api/v1
func main() {
	// log init
	logHandler := slog.NewJSONHandler(os.Stdout, nil)
	slog.SetDefault(slog.New(logHandler))

	// DB Connection, return handler
	slog.Info("Start DB Connection")
	noticeHdl, documentHdl, contactHdl := dbs.DBInit()
	slog.Info("DB Connection Success")

	// Gin init
	slog.Info("Gin Init")

	// v1 Group
	mainRounter := gin.Default()

	// health check
	mainRounter.GET("/health", func(c *gin.Context) {
		// DB Connection Check
		database, _ := dbs.DB.DB()
		if err := database.Ping(); err != nil {
			c.JSON(500, gin.H{"status": "unhealth", "reason": "db connection failed"})
			return
		}

		c.JSON(200, gin.H{"status": "ok"})
	})

	v1 := mainRounter.Group("/api/v1/client")
	slog.Info("Define Routes : version 1")
	{
		// Notice
		v1.GET("/notices", noticeHdl.GetAllNotices)
		v1.GET("/notices/:id", noticeHdl.GetNoticeByID)

		// Document
		v1.GET("/documents", documentHdl.GetAllDocuments)
		v1.GET("/documents/:id", documentHdl.GetDocumentsByID)

		// Contact
		v1.POST("/contacts", contactHdl.CreateContact)
	}

	// Server Init
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = ":8080"
	}
	slog.Info("Check Environment Variable and port : ", "SERVER_PORT", port)

	if err := mainRounter.Run(port); err != nil {
		slog.Error("Gin Server failed to start", "error", err)
	}
}
