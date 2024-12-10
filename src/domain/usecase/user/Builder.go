package userUseCase

import (
	repositoryInterface "task-tracker/boundary/repository"
	"task-tracker/infrastructure/errors"
	jwtServiceInterface "task-tracker/infrastructure/security/jwtService/interface"
)

type Builder struct {
	userUseCase *UserUseCase
}

func NewBuilder() *Builder {
	return &Builder{
		userUseCase: &UserUseCase{},
	}
}

func (b *Builder) JwtService(jwtService jwtServiceInterface.JWTService) *Builder {
	b.userUseCase.jwtService = jwtService
	return b
}

func (b *Builder) UserRepo(userRepo repositoryInterface.UserRepository) *Builder {
	b.userUseCase.userRepo = userRepo
	return b
}

func (b *Builder) Build() (*UserUseCase, error) {
	err := b.checkRequiredFields()
	if err != nil {
		return nil, err
	}
	return b.userUseCase, nil
}

func (b *Builder) checkRequiredFields() error {
	if b.userUseCase.jwtService == nil {
		return errors.ErrJWTServiceIsRequired
	}
	if b.userUseCase.userRepo == nil {
		return errors.ErrRepositoryIsRequired
	}
	return nil
}
