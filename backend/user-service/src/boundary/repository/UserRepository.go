package repositoryInterface

import (
	"context"
	emailPrimitive "github.com/revazcus/task-tracker/backend/common/domainPrimitive/email"
	idPrimitive "github.com/revazcus/task-tracker/backend/common/domainPrimitive/id"
	userEntity "github.com/revazcus/task-tracker/backend/user-service/domain/entity"
	passwordPrimitive "github.com/revazcus/task-tracker/backend/user-service/domain/entity/password"
	usernamePrimitive "github.com/revazcus/task-tracker/backend/user-service/domain/entity/username"
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
