package userDto

import userEntity "user-service/domain/entity"

type UserResponseDto struct {
	User  *userEntity.User
	Token string
}
