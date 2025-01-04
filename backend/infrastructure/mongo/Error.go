package mongoRepo

import "github.com/revazcus/task-tracker/backend/infrastructure/errors"

var (
	ErrMongoDbIsRequired = errors.NewError("SYS", "MongoDB is required")
	ErrLoggerIsRequired  = errors.NewError("SYS", "Logger is required")
)
