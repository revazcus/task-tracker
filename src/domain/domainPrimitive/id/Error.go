package idPrimitive

import "errors"

var (
	ErrEntityIdIsEmpty   = errors.New("id is empty")
	ErrEntityIdIsInvalid = errors.New("invalid id")
)
