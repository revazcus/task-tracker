package usecase

import "task-tracker/boundary/dto"

type PermissionUseCaseInterface interface {
	GetById(id int) (*dto.PermissionDto, error)
	Create(createDto *dto.PermissionDto) (*dto.PermissionDto, error)
	Update(userDto *dto.PermissionDto) (*dto.PermissionDto, error)
	Delete(id int) error
}
