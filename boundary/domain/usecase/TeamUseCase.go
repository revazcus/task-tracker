package usecase

import "task-tracker/boundary/dto"

type TeamUseCaseInterface interface {
	GetById(id int) (*dto.TeamDto, error)
	Create(createDto *dto.TeamDto) (*dto.TeamDto, error)
	Update(userDto *dto.TeamDto) (*dto.TeamDto, error)
	Delete(id int) error
}
