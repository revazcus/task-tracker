package emailPrimitive

import "errors"

var (
	ErrEmailIsEmpty   = errors.New("email is empty")
	ErrEmailIsInvalid = errors.New("invalid email")
)
