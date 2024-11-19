package usecase

import (
	"errors"
	"task-tracker/boundary/dto"
)

// UserUseCase имплементирует интерфейс UserUseCaseInterface через реализацию методов
type UserUseCase struct {
}

func (u UserUseCase) Create(createDto *dto.CreateUserDto) (*dto.CreateUserDto, error) {
	if createDto.Username == "" {
		return nil, errors.New("username is empty")
	}
	if createDto.Email == "" {
		return nil, errors.New("email is empty")
	}
	if createDto.Password == "" {
		return nil, errors.New("password is empty")
	}
	return createDto, nil // заглушка с возвратом входящих данных
}

// TODO докинуть GET | UPDATE | DELETE
