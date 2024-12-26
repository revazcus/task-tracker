package repository

import "github.com/stretchr/testify/mock"

type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) DeleteUser(userId int) error {
	args := m.Called(userId)
	return args.Error(0)
}
