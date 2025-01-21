package userRest

import (
	"infrastructure/errors"
	loggerInterface "infrastructure/logger/interface"
	restServerController "infrastructure/restServer/controller"
	"user-service/boundary/domain/usecase"
)

type Builder struct {
	controller *UserController
	errors     *errors.Errors
}

func NewBuilder() *Builder {
	return &Builder{
		controller: &UserController{},
		errors:     errors.NewErrors(),
	}
}

func (b *Builder) Logger(logger loggerInterface.Logger) *Builder {
	b.controller.logger = logger
	return b
}

func (b *Builder) BaseController(baseController *restServerController.BaseController) *Builder {
	b.controller.BaseController = baseController
	return b
}

func (b *Builder) UserUseCase(userUseCase usecase.UserUseCaseInterface) *Builder {
	b.controller.userUseCase = userUseCase
	return b
}

func (b *Builder) Build() (*UserController, error) {
	b.checkRequiredFields()
	if b.errors.IsPresent() {
		return nil, b.errors
	}
	b.fillDefaultFields()
	return b.controller, nil
}

func (b *Builder) checkRequiredFields() {
	if b.controller.logger == nil {
		b.errors.AddError(errors.NewError("SYS", "UserControllerBuilder: Logger is required"))
	}
	if b.controller.BaseController == nil {
		b.errors.AddError(errors.NewError("SYS", "UserControllerBuilder: BaseController is required"))
	}
	if b.controller.userUseCase == nil {
		b.errors.AddError(errors.NewError("SYS", "UserControllerBuilder: UserUseCase is required"))
	}
}

func (b *Builder) fillDefaultFields() {

}
