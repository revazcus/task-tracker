package usecase

import (
	"task-tracker/boundary/dto"
)

type TaskUseCase struct {
}

func (t TaskUseCase) GetById(id int) (*dto.TaskDto, error) {
	//TODO implement me
	panic("implement me")
}

func (t TaskUseCase) Create(createDto *dto.TaskDto) (*dto.TaskDto, error) {
	//TODO implement me
	panic("implement me")
}

func (t TaskUseCase) Update(userDto *dto.TaskDto) (*dto.TaskDto, error) {
	//TODO implement me
	panic("implement me")
}

func (t TaskUseCase) Delete(id int) error {
	//TODO implement me
	panic("implement me")
}
