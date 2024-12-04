package usecase

import "task-tracker/boundary/dto"

type TeamUseCaseInterface interface {
	GetById(id int) (*dto.TeamDto, error)
	Create(dto *dto.TeamDto) (*dto.TeamDto, error)
	Update(dto *dto.TeamDto) (*dto.TeamDto, error)
	Delete(id int) error
}
