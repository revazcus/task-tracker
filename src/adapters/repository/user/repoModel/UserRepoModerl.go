package repoModel

import (
	emailPrimitive "task-tracker/domain/domainPrimitive/email"
	idPrimitive "task-tracker/domain/domainPrimitive/id"
	passwordPrimitive "task-tracker/domain/domainPrimitive/password"
	userEntity "task-tracker/domain/entity/user"
)

type UserRepoModel struct {
	Id       string `bson:"user_id"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
}

func NewUserRepoModel(user *userEntity.User) *UserRepoModel {
	// TODO сейчас генерация ID на стороне монги, поэтому добавлена проверка для кейсов с созданием моделей
	var id string
	if user.ID() != nil {
		id = string(*user.ID())
	}
	return &UserRepoModel{
		Id:       id,
		Email:    string(*user.Email()),
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

	password, err := passwordPrimitive.PasswordFrom(m.Password)
	if err != nil {
		return nil, err
	}
	return userEntity.NewBuilder().
		Id(&id).
		Email(&email).
		Password(password).
		Build()
}
