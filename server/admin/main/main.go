package main

import (
	"os"

	"log/slog"

	"github.com/gin-gonic/gin"

	dbs "github.com/ahmola/Hanwoo-Tax-Corporation/pkg/common/dbs"
)

// @title				User API
// @version				2.0
// @description			API built with Gin
// @host				localhost:8081
// @BasePath			/api/v2
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
	r := gin.Default()

	// v1 Group
	v1 := r.Group("/api/v1/admin")
	slog.Info("Define Routes : version 1")
	{
		// Notice
		v1.POST("/notice", noticeHdl.CreateNotice)

		// Document
		v1.POST("/document", documentHdl.CreateDocument)

		// Contact
		v1.GET("/contact", contactHdl.GetCotectsByID)
	}

	// Server Init
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = ":8080"
	}
	slog.Info("Check Environment Variable and port : ", "SERVER_PORT", port)

	if err := r.Run(port); err != nil {
		slog.Error("Gin Server failed to start", "error", err)
	}
}
