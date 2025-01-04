package profileRepoModel

import profilePrimitive "github.com/revazcus/task-tracker/backend/common/domainPrimitive/profile"

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
	profile, err := profilePrimitive.NewProfile(m.FirstName, m.LastName)
	if err != nil {
		return nil, err
	}

	return profile, nil
}
