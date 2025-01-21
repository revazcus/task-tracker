package userEntity

import "infrastructure/errors"

var (
	ErrProfileIsRequired   = errors.NewError("USE-001", "Профиль должен быть заполнен")
	ErrEmailIsRequired     = errors.NewError("USE-002", "Email должен быть заполнен")
	ErrUsernameIsRequired  = errors.NewError("USE-003", "Логин должен быть заполнен")
	ErrPasswordIsRequired  = errors.NewError("USE-004", "Пароль должен быть заполнен")
	ErrRoleIsRequired      = errors.NewError("USE-005", "Роль должна быть заполнена")
	ErrAgreementIsRequired = errors.NewError("USE_006", "Согласие должно быть заполнено")

	ErrInvalidUsernameOrPassword = errors.NewError("USE-005", "Неверный логин или пароль")
)
