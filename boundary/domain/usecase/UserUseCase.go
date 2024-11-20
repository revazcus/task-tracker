package usecaseInterface

import (
	"task-tracker/boundary/dto"
)

// UserUseCaseInterface в первичной реализации интерфейса методы на вход принимают только dto без контекста
type UserUseCaseInterface interface {
	GetById(id int) (*dto.UserDto, error)
	Create(createDto *dto.UserDto) (*dto.UserDto, error)
	Update(userDto *dto.UserDto) (*dto.UserDto, error)
	Delete(id int) error
}
