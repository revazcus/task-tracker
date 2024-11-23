package usecase

import "task-tracker/boundary/dto"

type ProjectUseCaseInterface interface {
	GetById(id int) (*dto.ProjectDto, error)
	Create(createDto *dto.ProjectDto) (*dto.ProjectDto, error)
	Update(userDto *dto.ProjectDto) (*dto.ProjectDto, error)
	Delete(id int) error
}
