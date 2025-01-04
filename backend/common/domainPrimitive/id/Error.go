package idPrimitive

import "github.com/revazcus/task-tracker/backend/infrastructure/errors"

var (
	ErrEntityIdIsEmpty   = errors.NewError("IDP-001", "Id is empty")
	ErrEntityIdIsInvalid = errors.NewError("IDP-002", "Invalid id")
)
