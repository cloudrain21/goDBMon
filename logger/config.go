package logger

import (
	log "github.com/sirupsen/logrus"
)

func Configure(logFile string) {
	writer := NewFileLogger(logFile)
	log.SetOutput(writer)
}
