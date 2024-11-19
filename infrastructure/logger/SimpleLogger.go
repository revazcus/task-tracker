package logger

import (
	"log"
	"os"
	"time"
)

type SimpleLogger struct {
	infoLogger  *log.Logger
	errorLogger *log.Logger
}

func NewSimpleLogger() *SimpleLogger {
	return &SimpleLogger{
		infoLogger:  log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		errorLogger: log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (r *SimpleLogger) LogInfo(method string, path string) {
	r.infoLogger.Printf("%s | Method: %s | Path: %s", time.Now().Format(time.RFC3339), method, path)
}

func (r *SimpleLogger) LogError(err error, method string, path string) {
	r.errorLogger.Printf("%s | Method: %s | Path: %s | Error: %v", time.Now().Format(time.RFC3339), method, path, err)
}
