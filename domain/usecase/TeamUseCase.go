package usecase

import (
	"task-tracker/boundary/dto"
)

type TeamUseCase struct {
}

func (t TeamUseCase) GetById(id int) (*dto.TeamDto, error) {
	//TODO implement me
	panic("implement me")
}

func (t TeamUseCase) Create(createDto *dto.TeamDto) (*dto.TeamDto, error) {
	//TODO implement me
	panic("implement me")
}

func (t TeamUseCase) Update(userDto *dto.TeamDto) (*dto.TeamDto, error) {
	//TODO implement me
	panic("implement me")
}

func (t TeamUseCase) Delete(id int) error {
	//TODO implement me
	panic("implement me")
}
