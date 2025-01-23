package restResponse

import (
	"fmt"
	"infrastructure/errors"
)

func ErrUnmarshalRequest(description string) error {
	errMsg := fmt.Sprintf("Malformer request. Cause - %s", description)
	return errors.NewError("SYS", errMsg)
}
