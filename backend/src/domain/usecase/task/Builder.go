package taskUseCase

import (
	repositoryInterface "task-tracker/boundary/repository"
	userUseCase "task-tracker/domain/usecase/user"
	"task-tracker/infrastructure/errors"
)

type Builder struct {
	taskUseCase *TaskUseCase
	errors      *errors.Errors
}

func NewBuilder() *Builder {
	return &Builder{
		taskUseCase: &TaskUseCase{},
		errors:      errors.NewErrors(),
	}
}

func (b *Builder) TaskRepo(taskRepo repositoryInterface.TaskRepository) *Builder {
	b.taskUseCase.taskRepo = taskRepo
	return b
}

// UserUseCase TODO переписать
func (b *Builder) UserUseCase(userUseCase *userUseCase.UserUseCase) *Builder {
	b.taskUseCase.userUseCase = userUseCase
	return b
}

func (b *Builder) Build() (*TaskUseCase, error) {
	b.checkRequiredFields()
	if b.errors.IsPresent() {
		return nil, b.errors
	}
	return b.taskUseCase, nil
}

func (b *Builder) checkRequiredFields() {
	if b.taskUseCase.taskRepo == nil {
		b.errors.AddError(errors.NewError("SYS", "TaskUseCaseBuilder: TaskRepository is required"))
	}
	if b.taskUseCase.userUseCase == nil {
		b.errors.AddError(errors.NewError("SYS", "TaskUseCaseBuilder: UserUseCase is required"))
	}
}
