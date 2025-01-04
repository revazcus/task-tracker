package taskTimeCosts

import "github.com/revazcus/task-tracker/backend/infrastructure/errors"

var (
	ErrInvalidMinutes = errors.NewError("TC-001", "Затраты времени должны быть больше 0")
)
