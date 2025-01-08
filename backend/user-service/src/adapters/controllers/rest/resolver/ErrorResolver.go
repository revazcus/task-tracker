package resolver

import (
	"infrastructure/errors"
	"net/http"
	userRepo "user-service/src/adapters/repository/user"
)

const (
	UnknownErrorCode = "UNKNOWN_CODE"
)

type ErrorResolver struct {
}

func NewErrorResolver() *ErrorResolver {
	return &ErrorResolver{}
}

func (er *ErrorResolver) GetErrorCode(err error) string {
	switch errType := err.(type) {
	case *errors.Error:
		return string(errType.Code())
	default:
		return UnknownErrorCode
	}
}

func (er *ErrorResolver) GetErrorText(err error) string {
	switch errType := err.(type) {
	case *errors.Error:
		return errType.Message()
	default:
		return err.Error()
	}
}

func (er *ErrorResolver) GetHttpCode(err error) int {
	errs, ok := err.(*errors.Error)
	if !ok {
		return http.StatusInternalServerError
	}

	switch errs.Code() {
	case userRepo.ErrUserNotFound.Code():
		return http.StatusNotFound
	default:
		return http.StatusBadRequest
	}
}
