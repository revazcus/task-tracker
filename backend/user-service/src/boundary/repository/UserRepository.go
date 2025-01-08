package repositoryInterface

import (
	emailPrimitive "common/domainPrimitive/email"
	idPrimitive "common/domainPrimitive/id"
	"context"
	userEntity "user-service/src/domain/entity"
	passwordPrimitive "user-service/src/domain/entity/password"
	usernamePrimitive "user-service/src/domain/entity/username"
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
