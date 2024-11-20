package usecase

import (
	"errors"
	"task-tracker/boundary/dto"
)

// UserUseCase имплементирует интерфейс UserUseCaseInterface через реализацию методов
type UserUseCase struct {
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

func (u UserUseCase) Create(userDto *dto.UserDto) (*dto.UserDto, error) {
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
	return userDto, nil
}

func (u UserUseCase) Update(userDto *dto.UserDto) (*dto.UserDto, error) {
	if userDto.Id <= 0 {
		return nil, errors.New("invalid id")
	}
	return userDto, nil // заглушка с возвратом входящих данных
}

func (u UserUseCase) Delete(id int) error {
	return nil
}
