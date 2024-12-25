package titlePrimitive

import "task-tracker/infrastructure/errors"

var (
	ErrTitleIsEmpty  = errors.NewError("TPR-001", "Заголовок пустой")
	ErrTitleIsToLong = errors.NewError("TPR-002", "Заголовок превышает 255 символов")
)
