package userRest

import (
	"task-tracker/boundary/domain/usecase"
	"task-tracker/infrastructure/errors"
	loggerInterface "task-tracker/infrastructure/logger/interface"
	restServerController "task-tracker/infrastructure/restServer/controller"
)

type UserControllerBuilder struct {
	controller *UserController
	errors     *errors.Errors
}

func NewBuilder() *UserControllerBuilder {
	return &UserControllerBuilder{
		controller: &UserController{},
		errors:     errors.NewErrors(),
	}
}

// Logger инициализирует поле logger в UserController
func (b *UserControllerBuilder) Logger(logger loggerInterface.Logger) *UserControllerBuilder {
	b.controller.logger = logger
	return b
}

// BaseController инициализирует поле baseController в UserController
func (b *UserControllerBuilder) BaseController(baseController *restServerController.BaseController) *UserControllerBuilder {
	b.controller.BaseController = baseController
	return b
}

// UserUseCase инициализирует поле userUseCase в UserController
func (b *UserControllerBuilder) UserUseCase(userUseCase usecase.UserUseCaseInterface) *UserControllerBuilder {
	b.controller.userUseCase = userUseCase
	return b
}

func (b *UserControllerBuilder) Build() (*UserController, error) {
	b.checkRequiredFields()
	if b.errors.IsPresent() {
		return nil, b.errors
	}
	b.fillDefaultFields()
	return b.controller, nil
}

func (b *UserControllerBuilder) checkRequiredFields() {
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

// Инициализирует дефолтные поля в UserController
func (b *UserControllerBuilder) fillDefaultFields() {

}
