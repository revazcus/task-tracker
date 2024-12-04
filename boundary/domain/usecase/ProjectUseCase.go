package usecase

import "task-tracker/boundary/dto"

type ProjectUseCaseInterface interface {
	GetById(id int) (*dto.ProjectDto, error)
	Create(dto *dto.ProjectDto) (*dto.ProjectDto, error)
	Update(dto *dto.ProjectDto) (*dto.ProjectDto, error)
	Delete(id int) error
}
