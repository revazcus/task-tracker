package passwordPrimitive

import "task-tracker/infrastructure/errors"

var (
	ErrPasswordLength = errors.NewError("SYS", "Пароль меньше 8 символов")
	ErrWrongPassword  = errors.NewError("SYS", "Wrong password")
)
