package descriptionPrimitive

import "task-tracker/infrastructure/errors"

var (
	ErrDescriptionIsEmpty = errors.NewError("DPR-001", "Описание пустое")
	ErrDescriptionToLong  = errors.NewError("DPR-002", "Описание превышает 500 символов")
)
