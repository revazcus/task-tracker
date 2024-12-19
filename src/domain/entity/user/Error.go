package userEntity

import "task-tracker/infrastructure/errors"

var (
	ErrEmailIsRequired           = errors.NewError("USE-001", "User email is required")
	ErrUsernameIsRequired        = errors.NewError("USE-002", "Username is required")
	ErrPasswordIsRequired        = errors.NewError("USE-003", "User password is required")
	ErrInvalidUsernameOrPassword = errors.NewError("USE-004", "Неверный логин или пароль")
)
