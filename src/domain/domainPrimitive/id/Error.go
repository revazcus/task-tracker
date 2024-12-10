package idPrimitive

import "errors"

var (
	ErrEntityIdIsEmpty     = errors.New("id is empty")
	ErrEntityIdWrongFormat = errors.New("id must be numeric")
	ErrEntityIdIsInvalid   = errors.New("id must be greater than zero")
)
