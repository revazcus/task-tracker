package shortUserRepoModel

import (
	userObject "common/domainObject/shortUser"
	emailPrimitive "common/domainPrimitive/email"
	idPrimitive "common/domainPrimitive/id"
	profileRepoModel "common/repoModel/profile"
)

type ShortUserRepoModel struct {
	Id      string                             `bson:"user_id"`
	Email   string                             `bson:"email"`
	Profile *profileRepoModel.ProfileRepoModel `bson:"profile"`
}

func ShortUserToRepoModel(shortUser *userObject.ShortUser) *ShortUserRepoModel {
	if shortUser == nil {
		return nil
	}
	return &ShortUserRepoModel{
		Id:      shortUser.ID().String(),
		Email:   shortUser.Email().String(),
		Profile: profileRepoModel.ProfileToRepoModel(shortUser.Profile()),
	}
}

func (m *ShortUserRepoModel) GetObject() (*userObject.ShortUser, error) {
	if m == nil {
		return nil, nil
	}
	userId, err := idPrimitive.EntityIdFrom(m.Id)
	if err != nil {
		return nil, err
	}

	email, err := emailPrimitive.EmailFrom(m.Email)
	if err != nil {
		return nil, err
	}

	profile, err := m.Profile.GetPrimitive()
	if err != nil {
		return nil, err
	}

	return userObject.NewShortUser(&userId, &email, profile), nil
}
