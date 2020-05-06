package logger

import (
	"log"
	"os"
)

var logger = GetLogger("./dbmon.log")

func GetLogger(logFile string) *log.Logger {
	f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	return log.New(f, "dbmon : ", log.Ldate|log.Ltime)
}
