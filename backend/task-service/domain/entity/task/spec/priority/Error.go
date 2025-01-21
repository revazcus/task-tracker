package taskPriority

import (
	"fmt"
	"infrastructure/errors"
)

func ErrUnsupportedPriority(unsupportedPriority string) *errors.Error {
	msg := fmt.Sprintf("Unsupported Priority = `%s`", unsupportedPriority)
	return errors.NewError("SYS", msg)
}
