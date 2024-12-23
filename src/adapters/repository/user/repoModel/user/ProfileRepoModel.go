package userRepoModel

import profilePrimitive "task-tracker/domain/entity/user/profile"

type ProfileRepoModel struct {
	FirstName string `bson:"first_name"`
	LastName  string `bson:"last_name"`
}

func ProfileToRepoModel(profile *profilePrimitive.Profile) *ProfileRepoModel {
	return &ProfileRepoModel{
		FirstName: profile.FirstName(),
		LastName:  profile.LastName(),
	}
}

func (m *ProfileRepoModel) GetPrimitive() (*profilePrimitive.Profile, error) {
	profile, err := profilePrimitive.NewBuilder().
		FirstName(m.FirstName).
		LastName(m.LastName).
		Build()
	if err != nil {
		return nil, err
	}

	return profile, nil
}
