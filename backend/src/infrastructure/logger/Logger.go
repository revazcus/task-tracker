package commonLogger

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"strings"
	loggerInterface "task-tracker/infrastructure/logger/interface"
	logModel "task-tracker/infrastructure/logger/model"
)

type Logger struct {
	logChan chan<- *logModel.LogData
}

func NewLogger(logChan chan<- *logModel.LogData) *Logger {
	return &Logger{logChan: logChan}
}

func (l *Logger) Error(ctx context.Context, err error, options ...logModel.Option) {
	opts := &logModel.Options{}
	for _, opt := range options {
		opt(opts)
	}
	l.logError(ctx, err, opts)
}

func (l *Logger) Errors(ctx context.Context, errs []error, options ...logModel.Option) {
	opts := &logModel.Options{}
	for _, opt := range options {
		opt(opts)
	}
	for _, err := range errs {
		l.logError(ctx, err, opts)
	}
}

func (l *Logger) Info(ctx context.Context, message string, options ...logModel.Option) {
	l.logMessage(ctx, logModel.InfoLvl, message, options...)
}

func (l *Logger) Warning(ctx context.Context, message string, options ...logModel.Option) {
	l.logMessage(ctx, logModel.WarnLvl, message, options...)
}

func (l *Logger) Debug(ctx context.Context, message string, options ...logModel.Option) {
	l.logMessage(ctx, logModel.DebugLvl, message, options...)
}

func (l *Logger) logError(ctx context.Context, err error, opts *logModel.Options) {
	extendedErr := errors.WithStack(err)
	logData := &logModel.LogData{
		Ctx:    ctx,
		Msg:    extendedErr.Error(),
		Fields: []*logModel.LogField{},
		Lvl:    logModel.ErrorLvl,
	}

	if opts.WithStackTrace() {
		var fileNames []string
		if stackTracerErr, ok := extendedErr.(loggerInterface.StackTracer); ok {
			stackTrace := stackTracerErr.StackTrace()
			if len(stackTrace) > 0 {
				for i := 1; i < len(stackTrace); i++ {
					fileNames = append(fileNames, fmt.Sprintf("%s:%d", stackTrace[i], stackTrace[i]))
				}
			}
		}
		logData.Fields = append(logData.Fields, &logModel.LogField{Key: logModel.FieldFilenameKey, String: strings.Join(fileNames, " <- ")})
	}

	if len(opts.GetFields()) > 0 {
		logData.Fields = append(logData.Fields, opts.GetFields()...)
	}
	if opts.GetComponent() != "" {
		logData.Fields = append(logData.Fields, &logModel.LogField{Key: logModel.FieldComponentKey, String: opts.GetComponent()})
	}

	go l.sendData(logData)
}

func (l *Logger) logMessage(ctx context.Context, lvl logModel.LogLvl, message string, options ...logModel.Option) {
	opts := &logModel.Options{}
	for _, opt := range options {
		opt(opts)
	}

	logMsg := logModel.NewLogMessage(lvl, message).
		SetComponent(opts.GetComponent()).
		SetFields(opts.GetFields()...)
	logData := &logModel.LogData{
		Ctx:    ctx,
		Msg:    logMsg.Message,
		Fields: append(logMsg.Fields, &logModel.LogField{Key: logModel.FieldComponentKey, String: logMsg.Component}),
		Lvl:    logMsg.Lvl,
	}

	go l.sendData(logData)
}

func (l *Logger) sendData(logData *logModel.LogData) {
	l.logChan <- logData
}
