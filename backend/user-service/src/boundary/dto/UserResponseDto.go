package userDto

import userEntity "user-service/src/domain/entity"

type UserResponseDto struct {
	User  *userEntity.User
	Token string
}
