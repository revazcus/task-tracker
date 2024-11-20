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

func (r *SimpleLogger) LogInfo(method string, path string, requestBody string) {
	r.infoLogger.Printf("%s | Method: %s | Path: %s | Body:\n%s", time.Now().Format(time.RFC3339), method, path, requestBody)
}

func (r *SimpleLogger) LogError(err error, method string, path string, requestBody string) {
	r.errorLogger.Printf("%s | Method: %s | Path: %s | Error: %v | Body:%s", time.Now().Format(time.RFC3339), method, path, err, requestBody)
}
