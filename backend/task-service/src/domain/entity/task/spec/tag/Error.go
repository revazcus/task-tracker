package tag

import (
	"fmt"
	"github.com/revazcus/task-tracker/backend/infrastructure/errors"
)

func ErrUnsupportedTag(unsupportedTag string) *errors.Error {
	msg := fmt.Sprintf("Unsupported Tag = `%s`", unsupportedTag)
	return errors.NewError("SYS", msg)
}
