package repositoryInterface

import (
	"context"
	userEntity "task-tracker/domain/entity/user"
)

type UserRepository interface {
	Init() error
	Create(ctx context.Context, user *userEntity.User) (string, error)
	Update(ctx context.Context, user *userEntity.User) error
	GetById(ctx context.Context, id string) (*userEntity.User, error)
	GetAll(ctx context.Context) ([]*userEntity.User, error)
	Delete(ctx context.Context, id string) error
}
