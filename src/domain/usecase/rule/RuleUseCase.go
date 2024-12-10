package ruleUseCase

import (
	"task-tracker/boundary/dto"
)

type RuleUseCase struct {
}

func (r RuleUseCase) GetById(id int) (*dto.RuleDto, error) {
	//TODO implement me
	panic("implement me")
}

func (r RuleUseCase) Create(createDto *dto.RuleDto) (*dto.RuleDto, error) {
	//TODO implement me
	panic("implement me")
}

func (r RuleUseCase) Update(userDto *dto.RuleDto) (*dto.RuleDto, error) {
	//TODO implement me
	panic("implement me")
}

func (r RuleUseCase) Delete(id int) error {
	//TODO implement me
	panic("implement me")
}
