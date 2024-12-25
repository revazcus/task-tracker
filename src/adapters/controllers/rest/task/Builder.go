package taskRest

import (
	"task-tracker/boundary/domain/usecase"
	"task-tracker/infrastructure/errors"
	loggerInterface "task-tracker/infrastructure/logger/interface"
	restServerController "task-tracker/infrastructure/restServer/controller"
)

type Builder struct {
	controller *TaskController
	errors     *errors.Errors
}

func NewBuilder() *Builder {
	return &Builder{
		controller: &TaskController{},
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

func (b *Builder) TaskUseCase(userUseCase usecase.TaskUseCaseInterface) *Builder {
	b.controller.taskUseCase = userUseCase
	return b
}

func (b *Builder) Build() (*TaskController, error) {
	b.checkRequiredFields()
	if b.errors.IsPresent() {
		return nil, b.errors
	}
	b.fillDefaultFields()
	return b.controller, nil
}

func (b *Builder) checkRequiredFields() {
	if b.controller.logger == nil {
		b.errors.AddError(errors.NewError("SYS", "TaskControllerBuilder: Logger is required"))
	}
	if b.controller.BaseController == nil {
		b.errors.AddError(errors.NewError("SYS", "TaskControllerBuilder: BaseController is required"))
	}
	if b.controller.taskUseCase == nil {
		b.errors.AddError(errors.NewError("SYS", "TaskControllerBuilder: TaskUseCase is required"))
	}
}

func (b *Builder) fillDefaultFields() {

}
