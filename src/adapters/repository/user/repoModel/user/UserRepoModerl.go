package userRepoModel

import (
	emailPrimitive "task-tracker/domain/domainPrimitive/email"
	idPrimitive "task-tracker/domain/domainPrimitive/id"
	passwordPrimitive "task-tracker/domain/domainPrimitive/password"
	usernamePrimitive "task-tracker/domain/domainPrimitive/username"
	userEntity "task-tracker/domain/entity/user"
)

type UserRepoModel struct {
	Id       string `bson:"user_id"`
	Email    string `bson:"email"`
	Username string `bson:"username"`
	Password string `bson:"password"`
}

func UserToRepoModel(user *userEntity.User) *UserRepoModel {
	return &UserRepoModel{
		Id:       string(*user.ID()),
		Email:    string(*user.Email()),
		Username: string(*user.Username()),
		Password: string(*user.Password()),
	}
}

func (m *UserRepoModel) GetEntity() (*userEntity.User, error) {
	id, err := idPrimitive.EntityIdFrom(m.Id)
	if err != nil {
		return nil, err
	}

	email, err := emailPrimitive.EmailFrom(m.Email)
	if err != nil {
		return nil, err
	}

	username, err := usernamePrimitive.UsernameFrom(m.Username)
	if err != nil {
		return nil, err
	}

	password := passwordPrimitive.Password(m.Password)

	return userEntity.NewBuilder().
		Id(&id).
		Email(&email).
		Username(&username).
		Password(&password).
		Build()
}
