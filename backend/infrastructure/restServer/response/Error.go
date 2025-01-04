package response

import (
	"fmt"
	"github.com/revazcus/task-tracker/backend/infrastructure/errors"
)

func ErrUnmarshalRequest(description string) error {
	errMsg := fmt.Sprintf("Malformer request. Cause - %s", description)
	return errors.NewError("SYS", errMsg)
}
