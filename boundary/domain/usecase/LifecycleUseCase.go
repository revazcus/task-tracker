package usecase

import "task-tracker/boundary/dto"

type LifecycleUseCaseInterface interface {
	GetById(id int) (*dto.LifecycleDto, error)
	Create(createDto *dto.LifecycleDto) (*dto.LifecycleDto, error)
	Update(userDto *dto.LifecycleDto) (*dto.LifecycleDto, error)
	Delete(id int) error
}
