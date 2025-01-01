package userUseCase

import (
	"context"
	userDto "task-tracker/boundary/dto/user"
	"task-tracker/boundary/repository"
	idPrimitive "task-tracker/common/domainPrimitive/id"
	userEntity "task-tracker/domain/entity/user"
	agreementPrimitive "task-tracker/domain/entity/user/agreement"
	emailPrimitive "task-tracker/domain/entity/user/email"
	passwordPrimitive "task-tracker/domain/entity/user/password"
	profilePrimitive "task-tracker/domain/entity/user/profile"
	"task-tracker/domain/entity/user/spec"
	usernamePrimitive "task-tracker/domain/entity/user/username"
	"task-tracker/infrastructure/security/jwtService"
	jwtServiceInterface "task-tracker/infrastructure/security/jwtService/interface"
	commonTime "task-tracker/infrastructure/tools/time"
)

type UserUseCase struct {
	userRepo   repositoryInterface.UserRepository
	jwtService jwtServiceInterface.JWTService
}

func (u UserUseCase) CreateUser(ctx context.Context, userCreateDto *userDto.UserDto) (*userDto.UserResponseDto, error) {
	agreement, err := agreementPrimitive.NewBuilder().
		Accepted(userCreateDto.Agreement).
		AcceptedDate(commonTime.Now()).
		Build()
	if err != nil {
		return nil, err
	}

	profile, err := profilePrimitive.NewBuilder().
		FirstName(userCreateDto.FirstName).
		LastName(userCreateDto.LastName).
		Build()
	if err != nil {
		return nil, err
	}

	email, err := emailPrimitive.EmailFrom(userCreateDto.Email)
	if err != nil {
		return nil, err
	}

	password, err := passwordPrimitive.PasswordFrom(userCreateDto.Password)
	if err != nil {
		return nil, err
	}

	username, err := usernamePrimitive.UsernameFrom(userCreateDto.Username)
	if err != nil {
		return nil, err
	}

	user, err := userEntity.NewBuilder().
		Profile(profile).
		Email(&email).
		Username(&username).
		Password(password).
		Role(spec.Roles.Admin()).
		Agreement(agreement).
		Build()
	if err != nil {
		return nil, err
	}

	if err := u.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}

	token, err := u.jwtService.CreateUserToken(
		user.ID().String(),
		map[string]string{jwtService.RoleTokenKey: spec.Roles.Admin().String()})
	if err != nil {
		return nil, err
	}

	return &userDto.UserResponseDto{User: user, Token: token}, nil
}

func (u UserUseCase) GetAllUsers(ctx context.Context) ([]*userEntity.User, error) {
	foundUsers, err := u.userRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return foundUsers, nil
}

func (u UserUseCase) GetUserById(ctx context.Context, id string) (*userEntity.User, error) {
	userId := idPrimitive.EntityId(id)
	foundUser, err := u.userRepo.GetById(ctx, &userId)
	if err != nil {
		return nil, err
	}
	return foundUser, nil
}

func (u UserUseCase) UpdateUser(ctx context.Context, dto *userDto.UserDto) (*userEntity.User, error) {
	userId, err := idPrimitive.EntityIdFrom(dto.Id)
	if err != nil {
		return nil, err
	}

	agreement, err := agreementPrimitive.NewBuilder().
		Accepted(dto.Agreement).
		AcceptedDate(commonTime.Now()).
		Build()
	if err != nil {
		return nil, err
	}

	profile, err := profilePrimitive.NewBuilder().
		FirstName(dto.FirstName).
		LastName(dto.LastName).
		Build()
	if err != nil {
		return nil, err
	}

	email, err := emailPrimitive.EmailFrom(dto.Email)
	if err != nil {
		return nil, err
	}

	password, err := passwordPrimitive.PasswordFrom(dto.Password)
	if err != nil {
		return nil, err
	}

	username, err := usernamePrimitive.UsernameFrom(dto.Username)
	if err != nil {
		return nil, err
	}

	// TODO подумать что делать с обязательными полями при обновлении юзера
	user, err := userEntity.NewBuilder().
		Id(&userId).
		Profile(profile).
		Email(&email).
		Username(&username).
		Password(password).
		Role(spec.Roles.Admin()).
		Agreement(agreement).
		Build()
	if err != nil {
		return nil, err
	}

	// TODO по факту мы обновляем только профиль
	updatedUser, err := u.userRepo.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (u UserUseCase) UpdateEmail(ctx context.Context, dto *userDto.UserDto) (*userEntity.User, error) {
	userId, err := idPrimitive.EntityIdFrom(dto.Id)
	if err != nil {
		return nil, err
	}

	email, err := emailPrimitive.EmailFrom(dto.Email)
	if err != nil {
		return nil, err
	}

	updatedUser, err := u.userRepo.UpdateEmail(ctx, &userId, &email)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (u UserUseCase) UpdateUsername(ctx context.Context, dto *userDto.UserDto) (*userEntity.User, error) {
	userId, err := idPrimitive.EntityIdFrom(dto.Id)
	if err != nil {
		return nil, err
	}

	username, err := usernamePrimitive.UsernameFrom(dto.Username)
	if err != nil {
		return nil, err
	}

	updatedUser, err := u.userRepo.UpdateUsername(ctx, &userId, &username)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (u UserUseCase) UpdatePassword(ctx context.Context, dto *userDto.UserDto) (*userEntity.User, error) {
	userId, err := idPrimitive.EntityIdFrom(dto.Id)
	if err != nil {
		return nil, err
	}

	password, err := passwordPrimitive.PasswordFrom(dto.Password)
	if err != nil {
		return nil, err
	}

	updatedUser, err := u.userRepo.UpdatePassword(ctx, &userId, password)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (u UserUseCase) DeleteUser(ctx context.Context, id string) error {
	userId := idPrimitive.EntityId(id)
	if err := u.userRepo.DeleteById(ctx, &userId); err != nil {
		return err
	}
	return nil
}

func (u UserUseCase) LoginUser(ctx context.Context, dto *userDto.UserDto) (*userDto.UserResponseDto, error) {
	username := usernamePrimitive.Username(dto.Username)
	foundUser, err := u.userRepo.GetByUsername(ctx, &username)
	if err != nil {
		return nil, err
	}

	if err := foundUser.VerifyUsernameAndPassword(dto.Username, dto.Password); err != nil {
		return nil, err
	}

	// Наполняем токен
	token, err := u.jwtService.CreateUserToken(
		string(*foundUser.ID()),
		map[string]string{
			jwtService.RoleTokenKey: "ADMIN",
		},
	)

	if err != nil {
		return nil, err
	}

	return &userDto.UserResponseDto{User: foundUser, Token: token}, nil
}
