package passwordPrimitive

import "errors"

var (
	ErrPasswordLength  = errors.New("password shorter than 8 characters")
	ErrPasswordIsWrong = errors.New("password is wrong")
)
