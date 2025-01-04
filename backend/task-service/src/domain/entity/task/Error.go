package taskEntity

import "github.com/revazcus/task-tracker/backend/infrastructure/errors"

var (
	ErrTitleIsRequired       = errors.NewError("TAE-001", "Заголовок должен быть заполнен")
	ErrDescriptionIsRequired = errors.NewError("TAE-002", "Описание должно быть заполнено")
	ErrCreatorIsRequired     = errors.NewError("TAE-003", "Создатель должен быть заполнен")
)
