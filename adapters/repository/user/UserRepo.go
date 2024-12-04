package user

import (
	"task-tracker/boundary/dto"
	loggerInterface "task-tracker/infrastructure/logger/interface"
)

type UserRepo struct {
	logger loggerInterface.Logger
}

func (u *UserRepo) Create(userDto *dto.UserDto) error {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepo) Update(userDto *dto.UserDto) error {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepo) GetById(userId int) (*dto.UserDto, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepo) Delete(userId int) error {
	//TODO implement me
	panic("implement me")
}
