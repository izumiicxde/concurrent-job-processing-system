package logger

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	"concurrent-job-processing-system/internal/config"
)

type Logger struct {
	*slog.Logger
	file *os.File
}

func New(cfg *config.Config) *Logger {
	// create the log folder
	if err := os.MkdirAll(filepath.Dir(cfg.LogPath), 0o755); err != nil {
		panic(err)
	}

	// open or create log file.
	file, err := os.OpenFile(cfg.LogPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0o644)
	if err != nil {
		panic(err)
	}

	writer := io.MultiWriter(file, os.Stdout)
	handler := slog.NewTextHandler(writer, nil)

	return &Logger{Logger: slog.New(handler), file: file}
}

func (logger *Logger) HTTPRequest(method string, path string, status int, duration time.Duration, remoteIP string) {
	logger.Info("HTTP REQUEST COMPLETED", "component", "http", "method", method, "path", path, "status", status, "duration", duration, "remoteIP", remoteIP)
}

func (logger *Logger) HTTPError(method string, path string, status int, remoteIP string, err error) {
	logger.Error("HTTP REQUEST ERROR", "component", "http", "method", method, "path", path, "status", status, "remoteIP", remoteIP, "error", err)
}

func (logger *Logger) Close() {
	if err := logger.file.Close(); err != nil {
		fmt.Println("Failed to close the log file, Error:", err)
	}
}
