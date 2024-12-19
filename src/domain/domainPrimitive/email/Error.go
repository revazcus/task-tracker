package emailPrimitive

import "task-tracker/infrastructure/errors"

var (
	ErrEmailIsEmpty   = errors.NewError("EP-001", "Email пустой")
	ErrEmailIsInvalid = errors.NewError("EP-002", "Invalid email")
)
