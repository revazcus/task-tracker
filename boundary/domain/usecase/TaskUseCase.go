package usecase

import "task-tracker/boundary/dto"

type TaskUseCaseInterface interface {
	GetById(id int) (*dto.TaskDto, error)
	Create(createDto *dto.TaskDto) (*dto.TaskDto, error)
	Update(userDto *dto.TaskDto) (*dto.TaskDto, error)
	Delete(id int) error
}
