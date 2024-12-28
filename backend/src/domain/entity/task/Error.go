package taskEntity

import "task-tracker/infrastructure/errors"

var (
	ErrTitleIsRequired       = errors.NewError("TAE-001", "Заголовок должен быть заполнен")
	ErrDescriptionIsRequired = errors.NewError("TAE-002", "Описание должно быть заполнено")
	ErrCreatorIdIsRequired   = errors.NewError("TAE-003", "Создатель должен быть заполнен")
)
