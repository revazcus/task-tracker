package usecase

import (
	"task-tracker/boundary/dto"
)

type PermissionUseCase struct {
}

func (p PermissionUseCase) GetById(id int) (*dto.PermissionDto, error) {
	//TODO implement me
	panic("implement me")
}

func (p PermissionUseCase) Create(createDto *dto.PermissionDto) (*dto.PermissionDto, error) {
	//TODO implement me
	panic("implement me")
}

func (p PermissionUseCase) Update(userDto *dto.PermissionDto) (*dto.PermissionDto, error) {
	//TODO implement me
	panic("implement me")
}

func (p PermissionUseCase) Delete(id int) error {
	//TODO implement me
	panic("implement me")
}
