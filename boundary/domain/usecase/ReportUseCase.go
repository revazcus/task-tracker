package usecase

import "task-tracker/boundary/dto"

type ReportUseCaseInterface interface {
	GetById(id int) (*dto.ReportDto, error)
	Create(createDto *dto.ReportDto) (*dto.ReportDto, error)
	Update(userDto *dto.ReportDto) (*dto.ReportDto, error)
	Delete(id int) error
}
