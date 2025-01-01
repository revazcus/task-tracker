package repositoryInterface

import (
	"context"
	idPrimitive "task-tracker/common/domainPrimitive/id"
	userEntity "task-tracker/domain/entity/user"
	emailPrimitive "task-tracker/domain/entity/user/email"
	passwordPrimitive "task-tracker/domain/entity/user/password"
	usernamePrimitive "task-tracker/domain/entity/user/username"
)

type UserRepository interface {
	Init(ctx context.Context) error

	Create(ctx context.Context, user *userEntity.User) error

	GetAll(ctx context.Context) ([]*userEntity.User, error)
	GetById(ctx context.Context, userId *idPrimitive.EntityId) (*userEntity.User, error)
	GetByUsername(ctx context.Context, username *usernamePrimitive.Username) (*userEntity.User, error)

	Update(ctx context.Context, user *userEntity.User) (*userEntity.User, error)
	UpdateEmail(ctx context.Context, userId *idPrimitive.EntityId, email *emailPrimitive.Email) (*userEntity.User, error)
	UpdateUsername(ctx context.Context, userId *idPrimitive.EntityId, username *usernamePrimitive.Username) (*userEntity.User, error)
	UpdatePassword(ctx context.Context, userId *idPrimitive.EntityId, password *passwordPrimitive.Password) (*userEntity.User, error)

	DeleteById(ctx context.Context, userId *idPrimitive.EntityId) error
}
