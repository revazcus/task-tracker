package taskRepo

import "infrastructure/errors"

var (
	ErrTaskNotFound = errors.NewError("TAR-001", "Task не существует")
)
