package spec

import (
	"fmt"
	"infrastructure/errors"
)

func ErrUnsupportedRole(unsupportedRole string) *errors.Error {
	msg := fmt.Sprintf("Unsupported Role = `%s`", unsupportedRole)
	return errors.NewError("SYS", msg)
}
