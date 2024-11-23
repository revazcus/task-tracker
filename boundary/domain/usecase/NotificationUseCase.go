package usecase

import "task-tracker/boundary/dto"

type NotificationUseCaseInterface interface {
	GetById(id int) (*dto.NotificationDto, error)
	Create(createDto *dto.NotificationDto) (*dto.NotificationDto, error)
	Update(userDto *dto.NotificationDto) (*dto.NotificationDto, error)
	Delete(id int) error
}
