package reportUseCase

import (
	"task-tracker/boundary/dto"
)

type ReportUseCase struct {
}

func (r ReportUseCase) GetById(id int) (*dto.ReportDto, error) {
	//TODO implement me
	panic("implement me")
}

func (r ReportUseCase) Create(createDto *dto.ReportDto) (*dto.ReportDto, error) {
	//TODO implement me
	panic("implement me")
}

func (r ReportUseCase) Update(userDto *dto.ReportDto) (*dto.ReportDto, error) {
	//TODO implement me
	panic("implement me")
}

func (r ReportUseCase) Delete(id int) error {
	//TODO implement me
	panic("implement me")
}
