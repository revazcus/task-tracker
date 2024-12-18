package usecase

import (
	"context"
	userDto "task-tracker/boundary/dto/user"
	userEntity "task-tracker/domain/entity/user"
)

type UserUseCaseInterface interface {
	CreateUser(ctx context.Context, userCreateDto *userDto.UserDto) (*userDto.UserResponseDto, error)

	UpdateUser(ctx context.Context, dto *userDto.UserDto) (*userEntity.User, error)
	UpdateUserEmail(ctx context.Context, dto *userDto.UserDto) (*userEntity.User, error)
	UpdateUserPassword(ctx context.Context, dto *userDto.UserDto) (*userEntity.User, error)

	GetUserById(ctx context.Context, id string) (*userEntity.User, error)

	DeleteUser(ctx context.Context, id string) error

	LoginUser(ctx context.Context, dto *userDto.UserDto) (*userDto.UserResponseDto, error)
}
