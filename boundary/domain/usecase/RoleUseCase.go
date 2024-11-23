package usecase

import "task-tracker/boundary/dto"

type RoleUseCaseInterface interface {
	GetById(id int) (*dto.RoleDto, error)
	Create(createDto *dto.RoleDto) (*dto.RoleDto, error)
	Update(userDto *dto.RoleDto) (*dto.RoleDto, error)
	Delete(id int) error
}
