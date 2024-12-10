package lifecycleUseCase

import (
	"task-tracker/boundary/dto"
)

type LifecycleUseCase struct {
}

func (l LifecycleUseCase) GetById(id int) (*dto.LifecycleDto, error) {
	//TODO implement me
	panic("implement me")
}

func (l LifecycleUseCase) Create(createDto *dto.LifecycleDto) (*dto.LifecycleDto, error) {
	//TODO implement me
	panic("implement me")
}

func (l LifecycleUseCase) Update(userDto *dto.LifecycleDto) (*dto.LifecycleDto, error) {
	//TODO implement me
	panic("implement me")
}

func (l LifecycleUseCase) Delete(id int) error {
	//TODO implement me
	panic("implement me")
}
