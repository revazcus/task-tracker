package idPrimitive

import "infrastructure/errors"

var (
	ErrEntityIdIsEmpty   = errors.NewError("IDP-001", "Id is empty")
	ErrEntityIdIsInvalid = errors.NewError("IDP-002", "Invalid id")
)
