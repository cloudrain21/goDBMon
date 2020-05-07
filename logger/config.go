package logger

import (
	log "github.com/sirupsen/logrus"
	"io"
)

func Configure(writer io.Writer) {
	log.SetOutput(writer)
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
}
