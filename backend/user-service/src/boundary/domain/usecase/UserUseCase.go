package usecase

import (
	"context"
	userDto "github.com/revazcus/task-tracker/backend/user-service/boundary/dto/user"
	userEntity "github.com/revazcus/task-tracker/backend/user-service/domain/entity/user"
)

type UserUseCaseInterface interface {
	CreateUser(ctx context.Context, userCreateDto *userDto.UserDto) (*userDto.UserResponseDto, error)

	GetAllUsers(ctx context.Context) ([]*userEntity.User, error)
	GetUserById(ctx context.Context, id string) (*userEntity.User, error)

	UpdateUser(ctx context.Context, dto *userDto.UserDto) (*userEntity.User, error)
	UpdateEmail(ctx context.Context, dto *userDto.UserDto) (*userEntity.User, error)
	UpdateUsername(ctx context.Context, dto *userDto.UserDto) (*userEntity.User, error)
	UpdatePassword(ctx context.Context, dto *userDto.UserDto) (*userEntity.User, error)

	DeleteUser(ctx context.Context, id string) error

	LoginUser(ctx context.Context, dto *userDto.UserDto) (*userDto.UserResponseDto, error)
}
