package userObject

import (
	emailPrimitive "github.com/revazcus/task-tracker/backend/common/domainPrimitive/email"
	idPrimitive "github.com/revazcus/task-tracker/backend/common/domainPrimitive/id"
	profilePrimitive "github.com/revazcus/task-tracker/backend/common/domainPrimitive/profile"
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
