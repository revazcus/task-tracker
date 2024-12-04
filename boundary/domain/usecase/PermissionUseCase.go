package usecase

import "task-tracker/boundary/dto"

type PermissionUseCaseInterface interface {
	GetById(id int) (*dto.PermissionDto, error)
	Create(dto *dto.PermissionDto) (*dto.PermissionDto, error)
	Update(dto *dto.PermissionDto) (*dto.PermissionDto, error)
	Delete(id int) error
}
