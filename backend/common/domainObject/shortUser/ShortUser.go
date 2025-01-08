package userObject

import (
	emailPrimitive "common/domainPrimitive/email"
	idPrimitive "common/domainPrimitive/id"
	profilePrimitive "common/domainPrimitive/profile"
)

type ShortUser struct {
	id      *idPrimitive.EntityId
	email   *emailPrimitive.Email
	profile *profilePrimitive.Profile
}

func NewShortUser(userId *idPrimitive.EntityId, email *emailPrimitive.Email, profile *profilePrimitive.Profile) *ShortUser {
	return &ShortUser{id: userId, email: email, profile: profile}
}

func (u *ShortUser) ID() *idPrimitive.EntityId {
	return u.id
}

func (u *ShortUser) Email() *emailPrimitive.Email {
	return u.email
}

func (u *ShortUser) Profile() *profilePrimitive.Profile {
	return u.profile
}
