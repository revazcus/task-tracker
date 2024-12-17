package emailPrimitive

import "task-tracker/infrastructure/errors"

var (
	ErrEmailIsEmpty   = errors.NewError("SYS", "Email is empty")
	ErrEmailIsInvalid = errors.NewError("SYS", "Invalid email")
)
