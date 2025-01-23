package restServerController

import (
	"infrastructure/errors"
	loggerInterface "infrastructure/logger/interface"
	restResponse "infrastructure/restServer/response"
)

type Builder struct {
	responseService *restResponse.ResponseService
	logger          loggerInterface.Logger

	errors *errors.Errors
}

func NewBaseControllerBuilder() *Builder {
	return &Builder{
		errors: errors.NewErrors(),
	}
}

func (b *Builder) ResponseService(responseService *restResponse.ResponseService) *Builder {
	b.responseService = responseService
	return b
}

func (b *Builder) Logger(logger loggerInterface.Logger) *Builder {
	b.logger = logger
	return b
}

func (b *Builder) Build() (*BaseController, error) {
	b.checkRequiredFields()
	if b.errors.IsPresent() {
		return nil, b.errors
	}
	return b.createFromBuilder(), nil
}

func (b *Builder) checkRequiredFields() {
	if b.responseService == nil {
		b.errors.AddError(errors.NewError("SYS", "BaseControllerBuilder: ResponseService is required"))
	}
	if b.logger == nil {
		b.errors.AddError(errors.NewError("SYS", "BaseControllerBuilder: Logger is required"))
	}
}

func (b *Builder) createFromBuilder() *BaseController {
	return &BaseController{
		responseService: b.responseService,
		logger:          b.logger,
	}
}
