package usecase

import "task-tracker/boundary/dto"

type RuleUseCaseInterface interface {
	GetById(id int) (*dto.RuleDto, error)
	Create(dto *dto.RuleDto) (*dto.RuleDto, error)
	Update(dto *dto.RuleDto) (*dto.RuleDto, error)
	Delete(id int) error
}
