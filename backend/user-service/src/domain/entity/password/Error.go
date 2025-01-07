package passwordPrimitive

import "github.com/revazcus/task-tracker/backend/infrastructure/errors"

var (
	ErrPasswordLength = errors.NewError("SYS", "Пароль меньше 8 символов")
	ErrWrongPassword  = errors.NewError("SYS", "Wrong password")
)
