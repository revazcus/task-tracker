package userUseCase

import (
	"context"
	"errors"
	"task-tracker/boundary/dto"
	"task-tracker/boundary/repository"
	emailPrimitive "task-tracker/domain/domainPrimitive/email"
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

func (u UserUseCase) GetById(ctx context.Context, id string) (*dto.UserDto, error) {

	foundUser, err := u.userRepo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	// TODO скорее всего надо переписать, лишний маппинг
	userDto := dto.UserDto{
		Id:    string(*foundUser.ID()),
		Email: string(*foundUser.Email()),
	}

	return &userDto, nil
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

	userId, err := u.userRepo.Create(ctx, user)
	if err != nil {
		return nil, err
	}

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

	// TODO скорее всего надо переписать, лишний маппинг
	responseDto := dto.UserDto{
		Id:    userId,
		Email: string(*user.Email()),
		Token: token,
	}

	return &responseDto, nil
}

func (u UserUseCase) UpdateUser(ctx context.Context, userDto *dto.UserDto) (*dto.UserDto, error) {
	if userDto.Id == "" {
		return nil, errors.New("invalid id")
	}
	return userDto, nil // заглушка с возвратом входящих данных
}

func (u UserUseCase) DeleteUser(ctx context.Context, userId string) error {
	err := u.userRepo.Delete(ctx, userId)
	if err != nil {
		return err
	}
	return nil
}

func (u UserUseCase) LoginUser(ctx context.Context, reqDto *dto.UserDto) (*dto.UserDto, error) {

	foundUser, err := u.userRepo.GetById(ctx, reqDto.Id)
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
