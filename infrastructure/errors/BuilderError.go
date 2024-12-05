package errors

import "errors"

var (
	ErrLoggerIsRequired         = errors.New("logger is required")
	ErrBaseControllerIsRequired = errors.New("baseController is required")
	ErrUseCaseIsRequired        = errors.New("useCase is required")
	ErrSecretIsRequired         = errors.New("secret is required")
)
