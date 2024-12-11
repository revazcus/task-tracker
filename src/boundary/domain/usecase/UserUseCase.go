package usecase

import (
	"context"
	"task-tracker/boundary/dto"
)

// UserUseCaseInterface в первичной реализации интерфейса методы на вход принимают только dto без контекста
type UserUseCaseInterface interface {
	CreateUser(ctx context.Context, dto *dto.UserDto) (*dto.UserDto, error)

	UpdateUser(ctx context.Context, dto *dto.UserDto) (*dto.UserDto, error)
	UpdateUserEmail(ctx context.Context, dto *dto.UserDto) (*dto.UserDto, error)
	UpdateUserPassword(ctx context.Context, dto *dto.UserDto) (*dto.UserDto, error)

	GetUserById(ctx context.Context, id string) (*dto.UserDto, error)

	DeleteUser(ctx context.Context, id string) error

	LoginUser(ctx context.Context, dto *dto.UserDto) (*dto.UserDto, error)
}
