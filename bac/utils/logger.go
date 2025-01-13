package utils

import (
	"log"
	"os"
)

type Logger struct {
	logger *log.Logger
	file   *os.File
}

func NewLogger() *Logger {
	file, err := os.Create("app.log")
	if err != nil {
		log.Fatalf("Failed to create log file: %v", err)
	}

	return &Logger{
		logger: log.New(file, "", log.LstdFlags),
		file:   file,
	}
}

func (l *Logger) Cleanup() {
	l.file.Close()
}

func (l *Logger) Info(v ...interface{}) {
	l.logger.Println(v...)
}

func (l *Logger) Fatal(v ...interface{}) {
	l.logger.Fatal(v...)
}
