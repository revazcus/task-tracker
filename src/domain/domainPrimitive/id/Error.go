package idPrimitive

import "task-tracker/infrastructure/errors"

var (
	ErrEntityIdIsEmpty   = errors.NewError("SYS", "Id is empty")
	ErrEntityIdIsInvalid = errors.NewError("SYS", "Invalid id")
)
