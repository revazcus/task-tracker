package usernamePrimitive

import "task-tracker/infrastructure/errors"

var (
	ErrUsernameIsEmpty   = errors.NewError("SYS", "Username пустой")
	ErrUsernameLength    = errors.NewError("SYS", "Username должен содержать от 3 до 24 символов")
	ErrUsernameIsInvalid = errors.NewError("SYS", "Username может содержать только латиницу и цифры")
)
