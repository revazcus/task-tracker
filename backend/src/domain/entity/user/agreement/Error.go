package agreementPrimitive

import "task-tracker/infrastructure/errors"

var (
	ErrNotAcceptedAgreement   = errors.NewError("AGP-001", "Требуется согласие с пользовательским соглашением")
	ErrAcceptedDateIsRequired = errors.NewError("AGP-002", "Дата согласия не должна быть пустой")
)
