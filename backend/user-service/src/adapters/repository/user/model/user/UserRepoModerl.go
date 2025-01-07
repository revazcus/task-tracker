package userRepoModel

import (
	emailPrimitive "github.com/revazcus/task-tracker/backend/common/domainPrimitive/email"
	idPrimitive "github.com/revazcus/task-tracker/backend/common/domainPrimitive/id"
	profileRepoModel "github.com/revazcus/task-tracker/backend/common/repoModel/profile"
	commonTime "github.com/revazcus/task-tracker/backend/infrastructure/tools/time"
	userEntity "github.com/revazcus/task-tracker/backend/user-service/domain/entity"
	passwordPrimitive "github.com/revazcus/task-tracker/backend/user-service/domain/entity/password"
	"github.com/revazcus/task-tracker/backend/user-service/domain/entity/spec"
	usernamePrimitive "github.com/revazcus/task-tracker/backend/user-service/domain/entity/username"
)

type UserRepoModel struct {
	Id        string                             `bson:"user_id"`
	Profile   *profileRepoModel.ProfileRepoModel `bson:"profile"`
	Role      string                             `bson:"role"`
	Email     string                             `bson:"email"`
	Username  string                             `bson:"username"`
	Password  string                             `bson:"password"`
	Agreement *AgreementRepoModel                `bson:"agreement"`
	CreatedAt int64                              `bson:"created_at"`
}

func UserToRepoModel(user *userEntity.User) *UserRepoModel {
	return &UserRepoModel{
		Id:        string(*user.ID()),
		Profile:   profileRepoModel.ProfileToRepoModel(user.Profile()),
		Role:      user.Role().String(),
		Email:     string(*user.Email()),
		Username:  string(*user.Username()),
		Password:  string(*user.Password()),
		Agreement: AgreementToRepoModel(user.Agreement()),
		CreatedAt: user.CreatedAt().UnixNano(),
	}
}

func (m *UserRepoModel) GetEntity() (*userEntity.User, error) {
	id, err := idPrimitive.EntityIdFrom(m.Id)
	if err != nil {
		return nil, err
	}

	profile, err := m.Profile.GetPrimitive()
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

	role, err := spec.Roles.Of(m.Role)
	if err != nil {
		return nil, err
	}

	agreement, err := m.Agreement.GetPrimitive()
	if err != nil {
		return nil, err
	}

	createdAt := commonTime.FromUnixNano(m.CreatedAt)

	return userEntity.NewBuilder().
		Id(&id).
		Profile(profile).
		Role(role).
		Email(&email).
		Username(&username).
		Password(&password).
		Agreement(agreement).
		CreatedAt(createdAt).
		Build()
}
