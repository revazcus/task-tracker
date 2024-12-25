package taskDuration

import "task-tracker/infrastructure/errors"

var (
	ErrInvalidDayValue    = errors.NewError("TD-001", "Невалидный день")
	ErrInvalidHourValue   = errors.NewError("TD-002", "Невалидный час")
	ErrInvalidMinuteValue = errors.NewError("TD-003", "Невалидная минута")
)
