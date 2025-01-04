package descriptionPrimitive

import "github.com/revazcus/task-tracker/backend/infrastructure/errors"

var (
	ErrDescriptionIsEmpty = errors.NewError("DPR-001", "Описание пустое")
	ErrDescriptionToLong  = errors.NewError("DPR-002", "Описание превышает 500 символов")
)
