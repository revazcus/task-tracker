package userDto

import userEntity "github.com/revazcus/task-tracker/backend/user-service/domain/entity/user"

type UserResponseDto struct {
	User  *userEntity.User
	Token string
}
