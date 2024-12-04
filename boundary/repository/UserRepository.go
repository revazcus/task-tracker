package repositoryInterface

import "task-tracker/boundary/dto"

// UserRepository TODO добавить в методы context
type UserRepository interface {
	Create(dto *dto.UserDto) error
	Update(dto *dto.UserDto) error
	GetById(id int) (*dto.UserDto, error)
	Delete(id int) error
}
