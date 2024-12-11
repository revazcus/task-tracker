package userUseCase

import (
	"context"
	"errors"
	"task-tracker/boundary/dto"
	"task-tracker/boundary/repository"
	emailPrimitive "task-tracker/domain/domainPrimitive/email"
	idPrimitive "task-tracker/domain/domainPrimitive/id"
	passwordPrimitive "task-tracker/domain/domainPrimitive/password"
	userEntity "task-tracker/domain/entity/user"
	"task-tracker/infrastructure/security/jwtService"
	jwtServiceInterface "task-tracker/infrastructure/security/jwtService/interface"
)

// UserUseCase имплементирует интерфейс UserUseCaseInterface через реализацию методов
type UserUseCase struct {
	userRepo   repositoryInterface.UserRepository
	jwtService jwtServiceInterface.JWTService
}

func (u UserUseCase) CreateUser(ctx context.Context, userDto *dto.UserDto) (*dto.UserDto, error) {
	emailPrim, err := emailPrimitive.EmailFrom(userDto.Email)
	if err != nil {
		return nil, err
	}
	passwordPrim, err := passwordPrimitive.PasswordFrom(userDto.Password)
	if err != nil {
		return nil, err
	}

	user, err := userEntity.NewBuilder().
		Email(&emailPrim).
		Password(passwordPrim).
		Build()

	if err := u.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}

	userId := user.ID().String()

	// Наполняем токен при создании пользователя
	token, err := u.jwtService.CreateUserToken(
		userId,
		map[string]string{
			jwtService.RoleTokenKey: "ADMIN",
		},
	)

	if err != nil {
		return nil, err
	}

	// TODO переписать
	responseDto := dto.UserDto{
		Id:    userId,
		Email: string(*user.Email()),
		Token: token,
	}

	return &responseDto, nil
}

func (u UserUseCase) UpdateUser(ctx context.Context, userDto *dto.UserDto) (*dto.UserDto, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserUseCase) UpdateUserEmail(ctx context.Context, userDto *dto.UserDto) (*dto.UserDto, error) {
	if userDto.Id == "" {
		return nil, errors.New("invalid id")
	}
	userId := idPrimitive.EntityId(userDto.Id)

	email, err := emailPrimitive.EmailFrom(userDto.Email)
	if err != nil {
		return nil, err
	}

	if err := u.userRepo.UpdateEmail(ctx, &userId, &email); err != nil {
		return nil, err
	}

	// TODO переписать
	updatedDto := dto.UserDto{
		Id:    userDto.Id,
		Email: userDto.Email,
	}

	return &updatedDto, nil
}

func (u UserUseCase) UpdateUserPassword(ctx context.Context, userDto *dto.UserDto) (*dto.UserDto, error) {
	if userDto.Id == "" {
		return nil, errors.New("invalid id")
	}

	userId := idPrimitive.EntityId(userDto.Id)

	password, err := passwordPrimitive.PasswordFrom(userDto.Password)
	if err != nil {
		return nil, err
	}

	if err := u.userRepo.UpdatePassword(ctx, &userId, password); err != nil {
		return nil, err
	}

	// TODO переписать
	updatedDto := dto.UserDto{
		Id:       userDto.Id,
		Email:    userDto.Email,
		Password: "Пароль обновили по кайфу",
	}

	return &updatedDto, nil
}

func (u UserUseCase) GetUserById(ctx context.Context, id string) (*dto.UserDto, error) {
	userId := idPrimitive.EntityId(id)

	foundUser, err := u.userRepo.GetById(ctx, &userId)
	if err != nil {
		return nil, err
	}

	// TODO переписать
	userDto := dto.UserDto{
		Id:    string(*foundUser.ID()),
		Email: string(*foundUser.Email()),
	}

	return &userDto, nil
}

func (u UserUseCase) DeleteUser(ctx context.Context, id string) error {
	userId := idPrimitive.EntityId(id)

	err := u.userRepo.DeleteById(ctx, &userId)
	if err != nil {
		return err
	}
	return nil
}

func (u UserUseCase) LoginUser(ctx context.Context, reqDto *dto.UserDto) (*dto.UserDto, error) {
	userId := idPrimitive.EntityId(reqDto.Id)
	foundUser, err := u.userRepo.GetById(ctx, &userId)
	if err != nil {
		return nil, err
	}

	if err := foundUser.VerifyEmailAndPassword(reqDto.Email, reqDto.Password); err != nil {
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

	// TODO скорее всего надо переписать, лишний маппинг
	userDto := dto.UserDto{
		Id:    string(*foundUser.ID()),
		Email: string(*foundUser.Email()),
		Token: token,
	}

	return &userDto, nil
}
