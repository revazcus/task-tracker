package repositoryInterface

import (
	"context"
	emailPrimitive "task-tracker/domain/domainPrimitive/email"
	idPrimitive "task-tracker/domain/domainPrimitive/id"
	passwordPrimitive "task-tracker/domain/domainPrimitive/password"
	usernamePrimitive "task-tracker/domain/domainPrimitive/username"
	userEntity "task-tracker/domain/entity/user"
)

type UserRepository interface {
	Init(ctx context.Context) error

	Create(ctx context.Context, user *userEntity.User) error

	GetById(ctx context.Context, userId *idPrimitive.EntityId) (*userEntity.User, error)
	GetByUsername(ctx context.Context, username *usernamePrimitive.Username) (*userEntity.User, error)

	Update(ctx context.Context, user *userEntity.User) (*userEntity.User, error)
	UpdateEmail(ctx context.Context, userId *idPrimitive.EntityId, email *emailPrimitive.Email) (*userEntity.User, error)
	UpdatePassword(ctx context.Context, userId *idPrimitive.EntityId, password *passwordPrimitive.Password) (*userEntity.User, error)

	DeleteById(ctx context.Context, userId *idPrimitive.EntityId) error
}
