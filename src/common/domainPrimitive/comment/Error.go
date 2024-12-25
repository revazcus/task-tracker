package commentPrimitive

import "task-tracker/infrastructure/errors"

var (
	ErrCommentIsEmpty  = errors.NewError("CPR-001", "Комментарий пустой")
	ErrCommentIsToLong = errors.NewError("CPR-002", "Комментарий превышает 500 символов")
)
