package userEntity

import "errors"

var (
	ErrEmailIsRequired    = errors.New("user email is required")
	ErrPasswordIsRequired = errors.New("user password is required")
)
