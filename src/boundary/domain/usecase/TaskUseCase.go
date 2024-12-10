package usecase

import "task-tracker/boundary/dto"

type TaskUseCaseInterface interface {
	GetById(id int) (*dto.TaskDto, error)
	Create(dto *dto.TaskDto) (*dto.TaskDto, error)
	Update(dto *dto.TaskDto) (*dto.TaskDto, error)
	Delete(id int) error
}
