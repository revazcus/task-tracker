package userDto

import userEntity "github.com/revazcus/task-tracker/backend/user-service/domain/entity"

type UserResponseDto struct {
	User  *userEntity.User
	Token string
}
