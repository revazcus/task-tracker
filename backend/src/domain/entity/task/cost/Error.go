package taskTimeCosts

import "task-tracker/infrastructure/errors"

var (
	ErrInvalidMinutes = errors.NewError("TC-001", "Затраты времени должны быть больше 0")
)
