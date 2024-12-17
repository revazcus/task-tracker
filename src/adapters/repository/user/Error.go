package userRepo

import "task-tracker/infrastructure/errors"

var (
	ErrEmailAlreadyExist         = errors.NewError("SYS", "Email уже существует")
	ErrUsernameAlreadyExist      = errors.NewError("SYS", "Username уже существует")
	ErrInvalidUsernameOrPassword = errors.NewError("SYS", "Неверный логин или пароль")
)
