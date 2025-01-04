package emailPrimitive

import "github.com/revazcus/task-tracker/backend/infrastructure/errors"

var (
	ErrEmailIsEmpty   = errors.NewError("EMP-001", "Email пустой")
	ErrEmailIsInvalid = errors.NewError("EMP-002", "Некорректный Email")
)
