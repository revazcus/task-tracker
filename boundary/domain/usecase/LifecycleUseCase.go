package usecase

import "task-tracker/boundary/dto"

type LifecycleUseCaseInterface interface {
	GetById(id int) (*dto.LifecycleDto, error)
	Create(dto *dto.LifecycleDto) (*dto.LifecycleDto, error)
	Update(dto *dto.LifecycleDto) (*dto.LifecycleDto, error)
	Delete(id int) error
}
