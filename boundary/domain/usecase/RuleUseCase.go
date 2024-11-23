package usecase

import "task-tracker/boundary/dto"

type RuleUseCaseInterface interface {
	GetById(id int) (*dto.RuleDto, error)
	Create(createDto *dto.RuleDto) (*dto.RuleDto, error)
	Update(userDto *dto.RuleDto) (*dto.RuleDto, error)
	Delete(id int) error
}
