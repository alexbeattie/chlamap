// internal/utils/logger.go
package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

type Logger struct {
	file *os.File
}

func NewLogger() *Logger {
	if err := os.MkdirAll("logs", 0755); err != nil {
		log.Fatal("Failed to create logs directory:", err)
	}

	logFile := filepath.Join("logs", fmt.Sprintf("app_%s.log", time.Now().Format("2006-01-02")))
	f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}

	log.SetOutput(f)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	return &Logger{file: f}
}

func (l *Logger) Cleanup() {
	if l.file != nil {
		l.file.Close()
	}
}

func (l *Logger) Info(v ...interface{}) {
	log.Println(v...)
}

func (l *Logger) Error(v ...interface{}) {
	log.Println("ERROR:", v)
}

func (l *Logger) Fatal(v ...interface{}) {
	log.Fatal(v...)
}
