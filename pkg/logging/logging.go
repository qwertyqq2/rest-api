package logging

import (
	"log"
	"os"
)

type Logger struct {
	warningLogger *log.Logger
	errorLogger   *log.Logger
	infoLogger    *log.Logger
}

var l *Logger

func init() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	l = &Logger{
		warningLogger: log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile),
		infoLogger:    log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		errorLogger:   log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func GetLogger() *Logger {
	return l
}

func (l *Logger) Info(msg string) {
	l.infoLogger.Println(msg)
	log.Println(msg)
}

func (l *Logger) Warning(msg string) {
	l.warningLogger.Println(msg)
	log.Println(msg)
}

func (l *Logger) Error(msg string) {
	l.errorLogger.Println(msg)
	log.Println(msg)
}
