package userEntity

import (
	emailPrimitive "task-tracker/domain/domainPrimitive/email"
	idPrimitive "task-tracker/domain/domainPrimitive/id"
	passwordPrimitive "task-tracker/domain/domainPrimitive/password"
	usernamePrimitive "task-tracker/domain/domainPrimitive/username"
)

type Builder struct {
	id       *idPrimitive.EntityId
	email    *emailPrimitive.Email
	username *usernamePrimitive.Username
	password *passwordPrimitive.Password
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) Id(id *idPrimitive.EntityId) *Builder {
	b.id = id
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

func (b *Builder) Build() (*User, error) {
	err := b.checkRequiredFields()
	if err != nil {
		return nil, err
	}

	b.fillDefaultFields()

	return b.createFromBuilder(), nil

}

func (b *Builder) checkRequiredFields() error {
	if b.email == nil {
		return ErrEmailIsRequired
	}
	if b.username == nil {
		return ErrUsernameIsRequired
	}
	if b.password == nil {
		return ErrPasswordIsRequired
	}
	return nil
}

func (b *Builder) fillDefaultFields() {
	if b.id == nil {
		entityId := idPrimitive.NewEntityId()
		b.id = &entityId
	}
}

func (b *Builder) createFromBuilder() *User {
	return &User{
		id:       b.id,
		email:    b.email,
		username: b.username,
		password: b.password,
	}
}
