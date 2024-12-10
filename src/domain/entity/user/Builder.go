package userEntity

import (
	emailPrimitive "task-tracker/domain/domainPrimitive/email"
	idPrimitive "task-tracker/domain/domainPrimitive/id"
	passwordPrimitive "task-tracker/domain/domainPrimitive/password"
)

type Builder struct {
	id       *idPrimitive.EntityId
	email    *emailPrimitive.Email
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
	if b.password == nil {
		return ErrPasswordIsRequired
	}
	return nil
}

func (b *Builder) fillDefaultFields() {

}

func (b *Builder) createFromBuilder() *User {
	return &User{
		id:       b.id,
		email:    b.email,
		password: b.password,
	}
}
