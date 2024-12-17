package passwordPrimitive

import "task-tracker/infrastructure/errors"

var (
	ErrPasswordLength = errors.NewError("SYS", "Password shorter than 8 characters")
	ErrWrongPassword  = errors.NewError("SYS", "Wrong password")
)
