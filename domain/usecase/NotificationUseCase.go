package usecase

import (
	"task-tracker/boundary/dto"
)

type NotificationUseCase struct {
}

func (n NotificationUseCase) GetById(id int) (*dto.NotificationDto, error) {
	//TODO implement me
	panic("implement me")
}

func (n NotificationUseCase) Create(createDto *dto.NotificationDto) (*dto.NotificationDto, error) {
	//TODO implement me
	panic("implement me")
}

func (n NotificationUseCase) Update(userDto *dto.NotificationDto) (*dto.NotificationDto, error) {
	//TODO implement me
	panic("implement me")
}

func (n NotificationUseCase) Delete(id int) error {
	//TODO implement me
	panic("implement me")
}
