package userEntity

import "task-tracker/infrastructure/errors"

var (
	ErrEmailIsRequired           = errors.NewError("SYS", "User email is required")
	ErrUsernameIsRequired        = errors.NewError("SYS", "Username is required")
	ErrPasswordIsRequired        = errors.NewError("SYS", "User password is required")
	ErrInvalidUsernameOrPassword = errors.NewError("SYS", "Неверный логин или пароль")
)
