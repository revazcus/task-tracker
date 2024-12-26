package idPrimitive

import "task-tracker/infrastructure/errors"

var (
	ErrEntityIdIsEmpty   = errors.NewError("IDP-001", "Id is empty")
	ErrEntityIdIsInvalid = errors.NewError("IDP-002", "Invalid id")
)
