package usecase

import (
	"errors"
	"task-tracker/boundary/dto"
	"task-tracker/boundary/repository"
	"task-tracker/infrastructure/security/jwtService"
	jwtServiceInterface "task-tracker/infrastructure/security/jwtService/interface"
)

// UserUseCase имплементирует интерфейс UserUseCaseInterface через реализацию методов
type UserUseCase struct {
	userRepo   repositoryInterface.UserRepository
	jwtService jwtServiceInterface.JWTService
}

func NewUserUseCase(userRepo repositoryInterface.UserRepository, jwtService jwtServiceInterface.JWTService) *UserUseCase {
	return &UserUseCase{
		userRepo:   userRepo,
		jwtService: jwtService,
	}
}

func (u UserUseCase) GetById(id int) (*dto.UserDto, error) {
	// заглушка с возвратом предустановленных данных
	userDto := dto.UserDto{
		Id:       id,
		Username: "Гранд-адмирал Залупкинс",
		Email:    "pro100vasya@narod.ru",
		Password: "1488",
	}

	return &userDto, nil
}

func (u UserUseCase) CreateUser(userDto *dto.UserDto) (*dto.UserDto, error) {
	if userDto.Username == "" {
		return nil, errors.New("username is empty")
	}
	if userDto.Email == "" {
		return nil, errors.New("email is empty")
	}
	if userDto.Password == "" {
		return nil, errors.New("password is empty")
	}

	// заглушка с возвратом входящих данных + id
	userDto.Id = 1

	// сохраняем пользователя в БД
	//err := u.userRepo.Create(userDto)
	//if err != nil {
	//	return nil, err
	//}

	token, err := u.jwtService.CreateUserToken(
		userDto.Id,
		map[string]string{
			jwtService.RoleTokenKey: "ADMIN",
		},
	)

	if err != nil {
		return nil, err
	}

	userDto.Token = token

	return userDto, nil
}

func (u UserUseCase) UpdateUser(userDto *dto.UserDto) (*dto.UserDto, error) {
	if userDto.Id <= 0 {
		return nil, errors.New("invalid id")
	}
	return userDto, nil // заглушка с возвратом входящих данных
}

func (u UserUseCase) DeleteUser(userId int) error {
	err := u.userRepo.Delete(userId)
	if err != nil {
		return err
	}
	return nil
}
