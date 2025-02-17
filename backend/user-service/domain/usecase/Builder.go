package userUseCase

import (
	"infrastructure/errors"
	jwtServiceInterface "infrastructure/security/jwtService/interface"
	repositoryInterface "user-service/boundary/repository"
)

type Builder struct {
	userUseCase *UserUseCase
	errors      *errors.Errors
}

func NewBuilder() *Builder {
	return &Builder{
		userUseCase: &UserUseCase{},
		errors:      errors.NewErrors(),
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
	b.checkRequiredFields()
	if b.errors.IsPresent() {
		return nil, b.errors
	}
	return b.userUseCase, nil
}

func (b *Builder) checkRequiredFields() {
	if b.userUseCase.jwtService == nil {
		b.errors.AddError(errors.NewError("SYS", "UserUseCaseBuilder: JwtService is required"))
	}
	if b.userUseCase.userRepo == nil {
		b.errors.AddError(errors.NewError("SYS", "UserUseCaseBuilder: UserRepository is required"))
	}
}
