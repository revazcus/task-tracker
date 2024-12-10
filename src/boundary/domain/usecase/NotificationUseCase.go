package usecase

import "task-tracker/boundary/dto"

type NotificationUseCaseInterface interface {
	GetById(id int) (*dto.NotificationDto, error)
	Create(dto *dto.NotificationDto) (*dto.NotificationDto, error)
	Update(dto *dto.NotificationDto) (*dto.NotificationDto, error)
	Delete(id int) error
}
