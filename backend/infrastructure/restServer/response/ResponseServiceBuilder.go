package restResponse

import (
	"infrastructure/errors"
	loggerInterface "infrastructure/logger/interface"
)

type ResponseServiceBuilder struct {
	responseService *ResponseService
	errors          *errors.Errors
}

func NewResponseServiceBuilder() *ResponseServiceBuilder {
	return &ResponseServiceBuilder{
		responseService: &ResponseService{},
		errors:          errors.NewErrors(),
	}
}

func (b *ResponseServiceBuilder) ErrorResponseService(errorResponseService *ErrorResponseService) *ResponseServiceBuilder {
	b.responseService.errorResponseService = errorResponseService
	return b
}

func (b *ResponseServiceBuilder) Logger(logger loggerInterface.Logger) *ResponseServiceBuilder {
	b.responseService.logger = logger
	return b
}

func (b *ResponseServiceBuilder) Build() (*ResponseService, error) {
	b.checkRequiredFields()
	if b.errors.IsPresent() {
		return nil, b.errors
	}
	return b.responseService, nil
}

func (b *ResponseServiceBuilder) checkRequiredFields() {
	if b.responseService.errorResponseService == nil {
		b.errors.AddError(errors.NewError("SYS", "ResponseServiceBuilder: Logger is required"))
	}
	if b.responseService.logger == nil {
		b.errors.AddError(errors.NewError("SYS", "ResponseServiceBuilder: Logger is required"))
	}
}
