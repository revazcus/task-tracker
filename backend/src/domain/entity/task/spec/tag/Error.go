package taskTag

import (
	"fmt"
	"task-tracker/infrastructure/errors"
)

func ErrUnsupportedTag(unsupportedTag string) *errors.Error {
	msg := fmt.Sprintf("Unsupported Tag = `%s`", unsupportedTag)
	return errors.NewError("SYS", msg)
}
