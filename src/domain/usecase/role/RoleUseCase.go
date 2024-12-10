package roleUseCase

import (
	"task-tracker/boundary/dto"
)

type RoleUseCase struct {
}

func (r RoleUseCase) GetById(id int) (*dto.RoleDto, error) {
	//TODO implement me
	panic("implement me")
}

func (r RoleUseCase) Create(createDto *dto.RoleDto) (*dto.RoleDto, error) {
	//TODO implement me
	panic("implement me")
}

func (r RoleUseCase) Update(userDto *dto.RoleDto) (*dto.RoleDto, error) {
	//TODO implement me
	panic("implement me")
}

func (r RoleUseCase) Delete(id int) error {
	//TODO implement me
	panic("implement me")
}
