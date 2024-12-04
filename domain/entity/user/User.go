package userEntity

import (
	emailPrimitive "task-tracker/domain/domainPrimitive/email"
	idPrimitive "task-tracker/domain/domainPrimitive/id"
	passwordPrimitive "task-tracker/domain/domainPrimitive/password"
)

type User struct {
	id       idPrimitive.EntityId
	email    emailPrimitive.Email
	password passwordPrimitive.Password
}

func (u *User) ID() idPrimitive.EntityId {
	return u.id
}

func (u *User) Email() emailPrimitive.Email {
	return u.email
}
