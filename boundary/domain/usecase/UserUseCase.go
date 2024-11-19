package usecaseInterface

import (
	"task-tracker/boundary/dto"
)

// UserUseCaseInterface в первичной реализации интерфейса методы на вход принимают только dto без контекста
type UserUseCaseInterface interface {
	Create(createDto *dto.CreateUserDto) (*dto.CreateUserDto, error)

	// TODO докинуть GET | UPDATE | DELETE
}
