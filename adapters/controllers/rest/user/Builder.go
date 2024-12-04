package userRest

import (
	"task-tracker/boundary/domain/usecase"
	"task-tracker/infrastructure/errors"
	loggerInterface "task-tracker/infrastructure/logger/interface"
	restServerController "task-tracker/infrastructure/restServer/controller"
)

type UserControllerBuilder struct {
	controller *UserController
}

func NewBuilder() *UserControllerBuilder {
	return &UserControllerBuilder{
		controller: &UserController{},
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
	err := b.checkRequiredFields()
	if err != nil {
		return nil, err
	}
	b.fillDefaultFields()
	return b.controller, nil
}

func (b *UserControllerBuilder) checkRequiredFields() error {
	if b.controller.logger == nil {
		return errors.ErrLoggerIsRequired
	}
	if b.controller.BaseController == nil {
		return errors.ErrBaseControllerIsRequired
	}
	if b.controller.userUseCase == nil {
		return errors.ErrUseCaseIsRequired
	}
	return nil
}

// Инициализирует дефолтные поля в UserController
func (b *UserControllerBuilder) fillDefaultFields() {

}
