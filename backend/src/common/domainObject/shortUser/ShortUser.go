package userObject

import (
	idPrimitive "task-tracker/common/domainPrimitive/id"
	emailPrimitive "task-tracker/domain/entity/user/email"
	profilePrimitive "task-tracker/domain/entity/user/profile"
)

type ShortUser struct {
	id      *idPrimitive.EntityId
	email   *emailPrimitive.Email
	profile *profilePrimitive.Profile
}

func NewShortUser(userId *idPrimitive.EntityId, email *emailPrimitive.Email, profile *profilePrimitive.Profile) (*ShortUser, error) {
	return &ShortUser{id: userId, email: email, profile: profile}, nil
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
