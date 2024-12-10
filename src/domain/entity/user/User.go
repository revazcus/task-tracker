package userEntity

import (
	"errors"
	emailPrimitive "task-tracker/domain/domainPrimitive/email"
	idPrimitive "task-tracker/domain/domainPrimitive/id"
	passwordPrimitive "task-tracker/domain/domainPrimitive/password"
)

type User struct {
	id       *idPrimitive.EntityId
	email    *emailPrimitive.Email
	password *passwordPrimitive.Password
}

func (u *User) ID() *idPrimitive.EntityId {
	return u.id
}

func (u *User) Email() *emailPrimitive.Email {
	return u.email
}

func (u *User) Password() *passwordPrimitive.Password {
	return u.password
}

func (u *User) VerifyEmailAndPassword(email, password string) error {
	if !u.email.Verify(email) && !u.password.Verify(password) {
		return errors.New("неверный email или пароль")
	}
	return nil
}
