package taskTimeCosts

import "infrastructure/errors"

var (
	ErrInvalidMinutes = errors.NewError("TC-001", "Затраты времени должны быть больше 0")
)
