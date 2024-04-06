package logger

import (
	"log"
	"os"
)

var (
	infoLog  *log.Logger
	errorLog *log.Logger
)

func init() { //nolint:gochecknoinits
	infoLog = log.New(os.Stdout,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	errorLog = log.New(os.Stderr,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func Info(v ...any) {
	infoLog.Println(v...)
}

func Error(v ...any) {
	errorLog.Println(v...)
}

func Fatal(v ...any) {
	errorLog.Fatal(v...)
}
