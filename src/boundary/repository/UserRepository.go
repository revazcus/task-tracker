package repositoryInterface

import (
	"context"
	emailPrimitive "task-tracker/domain/domainPrimitive/email"
	idPrimitive "task-tracker/domain/domainPrimitive/id"
	passwordPrimitive "task-tracker/domain/domainPrimitive/password"
	userEntity "task-tracker/domain/entity/user"
)

type UserRepository interface {
	Create(ctx context.Context, user *userEntity.User) error

	Update(ctx context.Context, user *userEntity.User) error
	UpdateEmail(ctx context.Context, userId *idPrimitive.EntityId, email *emailPrimitive.Email) error
	UpdatePassword(ctx context.Context, userId *idPrimitive.EntityId, password *passwordPrimitive.Password) error

	GetById(ctx context.Context, userId *idPrimitive.EntityId) (*userEntity.User, error)

	DeleteById(ctx context.Context, userId *idPrimitive.EntityId) error
}
