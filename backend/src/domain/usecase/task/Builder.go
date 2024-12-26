package taskUseCase

import (
	repositoryInterface "task-tracker/boundary/repository"
	"task-tracker/infrastructure/errors"
)

type Builder struct {
	taskUserCase *TaskUseCase
	errors       *errors.Errors
}

func NewBuilder() *Builder {
	return &Builder{
		taskUserCase: &TaskUseCase{},
		errors:       errors.NewErrors(),
	}
}

func (b *Builder) TaskRepo(userRepo repositoryInterface.TaskRepository) *Builder {
	b.taskUserCase.taskRepo = userRepo
	return b
}

func (b *Builder) Build() (*TaskUseCase, error) {
	b.checkRequiredFields()
	if b.errors.IsPresent() {
		return nil, b.errors
	}
	return b.taskUserCase, nil
}

func (b *Builder) checkRequiredFields() {
	if b.taskUserCase.taskRepo == nil {
		b.errors.AddError(errors.NewError("SYS", "TaskUseCaseBuilder: TaskRepository is required"))
	}
}
