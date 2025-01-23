package restResponse

import (
	"infrastructure/errors"
	loggerInterface "infrastructure/logger/interface"
	restServerInterface "infrastructure/restServer/interface"
)

type ErrorResponseServiceBuilder struct {
	errorResponseService *ErrorResponseService
	errors               *errors.Errors
}

func NewErrorResponseServiceBuilder() *ErrorResponseServiceBuilder {
	return &ErrorResponseServiceBuilder{
		errorResponseService: &ErrorResponseService{},
		errors:               errors.NewErrors(),
	}
}

func (b *ErrorResponseServiceBuilder) ErrorResolver(errorResolver restServerInterface.ErrorResolver) *ErrorResponseServiceBuilder {
	b.errorResponseService.errorResolver = errorResolver
	return b
}

func (b *ErrorResponseServiceBuilder) Logger(logger loggerInterface.Logger) *ErrorResponseServiceBuilder {
	b.errorResponseService.logger = logger
	return b
}

func (b *ErrorResponseServiceBuilder) Build() (*ErrorResponseService, error) {
	b.checkRequiredFields()
	if b.errors.IsPresent() {
		return nil, b.errors
	}
	return b.errorResponseService, nil
}

func (b *ErrorResponseServiceBuilder) checkRequiredFields() {
	if b.errorResponseService.errorResolver == nil {
		b.errors.AddError(errors.NewError("SYS", "ErrorResponseServiceBuilder: ErrorResolver is required"))
	}
	if b.errorResponseService.logger == nil {
		b.errors.AddError(errors.NewError("SYS", "ErrorResponseServiceBuilder: Logger is required"))
	}
}
