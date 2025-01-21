package assessmentPrimitive

import "infrastructure/errors"

var (
	ErrInvalidAssessment = errors.NewError("APR-001", "Оценка времени должна быть больше 0")
)
