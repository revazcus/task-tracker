package taskStatus

import (
	"fmt"
	"infrastructure/errors"
)

func ErrUnsupportedStatus(unsupportedStatus string) *errors.Error {
	msg := fmt.Sprintf("Unsupported Status = `%s`", unsupportedStatus)
	return errors.NewError("SYS", msg)
}
