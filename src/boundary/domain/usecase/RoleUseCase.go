package usecase

import "task-tracker/boundary/dto"

type RoleUseCaseInterface interface {
	GetById(id int) (*dto.RoleDto, error)
	Create(dto *dto.RoleDto) (*dto.RoleDto, error)
	Update(dto *dto.RoleDto) (*dto.RoleDto, error)
	Delete(id int) error
}
