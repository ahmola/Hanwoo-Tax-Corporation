package main

import (
	"os"

	"log/slog"

	"github.com/gin-gonic/gin"

	dbs "github.com/ahmola/Hanwoo-Tax-Corporation/pkg/common/dbs"
)

// @title				Admin API
// @version				1.0
// @description			API built with Gin
// @host				localhost:8081
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

	v1 := mainRounter.Group("/api/v1/admin")
	slog.Info("Define Routes : version 1")
	{
		// Notice
		v1.POST("/notices", noticeHdl.CreateNotice)
		v1.DELETE("/notices/:id", noticeHdl.DeleteNotice)

		// Document
		v1.POST("/documents", documentHdl.CreateDocument)
		v1.DELETE("documents", documentHdl.DeleteDocument)

		// Contact
		v1.GET("/contacts/all", contactHdl.GetAllContacts)
		v1.GET("/contacts", contactHdl.GetCotactsByID)
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
