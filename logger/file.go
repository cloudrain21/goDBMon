package logger

import (
	"io"
	"log"
	"os"
)

type FileLogger struct {
	fileName string
	w        io.Writer
}

func NewFileLogger(fileName string) io.Writer {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	return &FileLogger{fileName, file}
}

func (f *FileLogger) Write(p []byte) (int, error) {
	return f.w.Write(p)
}
