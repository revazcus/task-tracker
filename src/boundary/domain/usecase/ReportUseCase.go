package usecase

import "task-tracker/boundary/dto"

type ReportUseCaseInterface interface {
	GetById(id int) (*dto.ReportDto, error)
	Create(dto *dto.ReportDto) (*dto.ReportDto, error)
	Update(dto *dto.ReportDto) (*dto.ReportDto, error)
	Delete(id int) error
}
