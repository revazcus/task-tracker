package usecase

import (
	"task-tracker/boundary/dto"
)

type ProjectUseCase struct {
}

func (p ProjectUseCase) GetById(id int) (*dto.ProjectDto, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProjectUseCase) Create(createDto *dto.ProjectDto) (*dto.ProjectDto, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProjectUseCase) Update(userDto *dto.ProjectDto) (*dto.ProjectDto, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProjectUseCase) Delete(id int) error {
	//TODO implement me
	panic("implement me")
}
