package userEntity

import (
	emailPrimitive "task-tracker/domain/domainPrimitive/email"
	idPrimitive "task-tracker/domain/domainPrimitive/id"
	passwordPrimitive "task-tracker/domain/domainPrimitive/password"
	usernamePrimitive "task-tracker/domain/domainPrimitive/username"
)

type User struct {
	id       *idPrimitive.EntityId
	username *usernamePrimitive.Username
	email    *emailPrimitive.Email
	password *passwordPrimitive.Password
}

func (u *User) ID() *idPrimitive.EntityId {
	return u.id
}

func (u *User) Email() *emailPrimitive.Email {
	return u.email
}

func (u *User) Username() *usernamePrimitive.Username {
	return u.username
}

func (u *User) Password() *passwordPrimitive.Password {
	return u.password
}

func (u *User) VerifyUsernameAndPassword(username, password string) error {
	if !u.username.Verify(username) || !u.password.Verify(password) {
		return ErrInvalidUsernameOrPassword
	}
	return nil
}
