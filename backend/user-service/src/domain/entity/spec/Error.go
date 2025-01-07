package spec

import (
	"fmt"
	"github.com/revazcus/task-tracker/backend/infrastructure/errors"
)

func ErrUnsupportedRole(unsupportedRole string) *errors.Error {
	msg := fmt.Sprintf("Unsupported Role = `%s`", unsupportedRole)
	return errors.NewError("SYS", msg)
}
