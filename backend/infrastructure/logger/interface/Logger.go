package loggerInterface

import (
	"context"
	logModel "github.com/revazcus/task-tracker/backend/infrastructure/logger/model"
)

type Logger interface {
	Error(ctx context.Context, err error, options ...logModel.Option)
	Errors(ctx context.Context, errs []error, options ...logModel.Option)
	Info(ctx context.Context, message string, options ...logModel.Option)
	Warning(ctx context.Context, message string, options ...logModel.Option)
	Debug(ctx context.Context, message string, options ...logModel.Option)
}
