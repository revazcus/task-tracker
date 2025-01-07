package userEntity

import (
	emailPrimitive "github.com/revazcus/task-tracker/backend/common/domainPrimitive/email"
	idPrimitive "github.com/revazcus/task-tracker/backend/common/domainPrimitive/id"
	profilePrimitive "github.com/revazcus/task-tracker/backend/common/domainPrimitive/profile"
	commonTime "github.com/revazcus/task-tracker/backend/infrastructure/tools/time"
	agreementPrimitive "github.com/revazcus/task-tracker/backend/user-service/domain/entity/agreement"
	passwordPrimitive "github.com/revazcus/task-tracker/backend/user-service/domain/entity/password"
	"github.com/revazcus/task-tracker/backend/user-service/domain/entity/spec"
	usernamePrimitive "github.com/revazcus/task-tracker/backend/user-service/domain/entity/username"
)

type User struct {
	id        *idPrimitive.EntityId
	profile   *profilePrimitive.Profile
	email     *emailPrimitive.Email
	username  *usernamePrimitive.Username
	password  *passwordPrimitive.Password
	role      spec.Role
	agreement *agreementPrimitive.Agreement
	createdAt *commonTime.Time
}

func (u *User) ID() *idPrimitive.EntityId {
	return u.id
}

func (u *User) Profile() *profilePrimitive.Profile {
	return u.profile
}

func (u *User) Username() *usernamePrimitive.Username {
	return u.username
}

func (u *User) Email() *emailPrimitive.Email {
	return u.email
}

func (u *User) Password() *passwordPrimitive.Password {
	return u.password
}

func (u *User) Role() spec.Role {
	return u.role
}

func (u *User) Agreement() *agreementPrimitive.Agreement {
	return u.agreement
}

func (u *User) CreatedAt() *commonTime.Time {
	return u.createdAt
}

func (u *User) VerifyUsernameAndPassword(username, password string) error {
	if !u.username.Verify(username) || !u.password.Verify(password) {
		return ErrInvalidUsernameOrPassword
	}
	return nil
}
