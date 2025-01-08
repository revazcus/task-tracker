package userRepo

import "infrastructure/errors"

var (
	ErrEmailAlreadyExist         = errors.NewError("USR-001", "Email уже существует")
	ErrUsernameAlreadyExist      = errors.NewError("USR-002", "Username уже существует")
	ErrInvalidUsernameOrPassword = errors.NewError("USR-003", "Неверный логин или пароль")
	ErrUserNotFound              = errors.NewError("USR-004", "User не существует")
)
