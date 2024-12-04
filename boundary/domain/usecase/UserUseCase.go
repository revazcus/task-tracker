package usecase

import (
	"task-tracker/boundary/dto"
)

// UserUseCaseInterface в первичной реализации интерфейса методы на вход принимают только dto без контекста
type UserUseCaseInterface interface {
	GetById(id int) (*dto.UserDto, error)
	CreateUser(dto *dto.UserDto) (*dto.UserDto, error)
	UpdateUser(dto *dto.UserDto) (*dto.UserDto, error)
	DeleteUser(id int) error
}
