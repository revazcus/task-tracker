package userEntity

import (
	emailPrimitive "github.com/revazcus/task-tracker/backend/common/domainPrimitive/email"
	idPrimitive "github.com/revazcus/task-tracker/backend/common/domainPrimitive/id"
	profilePrimitive "github.com/revazcus/task-tracker/backend/common/domainPrimitive/profile"
	"github.com/revazcus/task-tracker/backend/infrastructure/errors"
	commonTime "github.com/revazcus/task-tracker/backend/infrastructure/tools/time"
	agreementPrimitive "github.com/revazcus/task-tracker/backend/user-service/domain/entity/agreement"
	passwordPrimitive "github.com/revazcus/task-tracker/backend/user-service/domain/entity/password"
	"github.com/revazcus/task-tracker/backend/user-service/domain/entity/spec"
	usernamePrimitive "github.com/revazcus/task-tracker/backend/user-service/domain/entity/username"
)

type Builder struct {
	id        *idPrimitive.EntityId
	profile   *profilePrimitive.Profile
	email     *emailPrimitive.Email
	username  *usernamePrimitive.Username
	password  *passwordPrimitive.Password
	role      spec.Role
	agreement *agreementPrimitive.Agreement
	createdAt *commonTime.Time
	errors    *errors.Errors
}

func NewBuilder() *Builder {
	return &Builder{
		errors: errors.NewErrors(),
	}
}

func (b *Builder) Id(id *idPrimitive.EntityId) *Builder {
	b.id = id
	return b
}

func (b *Builder) Profile(profile *profilePrimitive.Profile) *Builder {
	b.profile = profile
	return b
}

func (b *Builder) Email(email *emailPrimitive.Email) *Builder {
	b.email = email
	return b
}

func (b *Builder) Username(username *usernamePrimitive.Username) *Builder {
	b.username = username
	return b
}

func (b *Builder) Password(password *passwordPrimitive.Password) *Builder {
	b.password = password
	return b
}

func (b *Builder) Role(role spec.Role) *Builder {
	b.role = role
	return b
}

func (b *Builder) Agreement(agreement *agreementPrimitive.Agreement) *Builder {
	b.agreement = agreement
	return b
}

func (b *Builder) CreatedAt(createdAt *commonTime.Time) *Builder {
	b.createdAt = createdAt
	return b
}

func (b *Builder) Build() (*User, error) {
	b.checkRequiredFields()
	if b.errors.IsPresent() {
		return nil, b.errors
	}

	b.fillDefaultFields()

	return b.createFromBuilder(), nil

}

func (b *Builder) checkRequiredFields() {
	if b.profile == nil {
		b.errors.AddError(ErrProfileIsRequired)
	}
	if b.email == nil {
		b.errors.AddError(ErrEmailIsRequired)
	}
	if b.username == nil {
		b.errors.AddError(ErrUsernameIsRequired)
	}
	if b.password == nil {
		b.errors.AddError(ErrPasswordIsRequired)
	}
	if b.role == "" {
		b.errors.AddError(ErrRoleIsRequired)
	}
	if b.agreement == nil {
		b.errors.AddError(ErrAgreementIsRequired)
	}
}

func (b *Builder) fillDefaultFields() {
	if b.id == nil {
		entityId := idPrimitive.NewEntityId()
		b.id = &entityId
	}
	if b.createdAt == nil {
		b.createdAt = commonTime.Now()
	}
}

func (b *Builder) createFromBuilder() *User {
	return &User{
		id:        b.id,
		profile:   b.profile,
		email:     b.email,
		username:  b.username,
		password:  b.password,
		role:      b.role,
		agreement: b.agreement,
		createdAt: b.createdAt,
	}
}
