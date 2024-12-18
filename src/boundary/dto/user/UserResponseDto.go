package userDto

import userEntity "task-tracker/domain/entity/user"

type UserResponseDto struct {
	User  *userEntity.User
	Token string
}
