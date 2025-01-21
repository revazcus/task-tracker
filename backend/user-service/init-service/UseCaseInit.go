package initService

import (
	"user-service/boundary/domain/usecase"
	userUseCase "user-service/domain/usecase"
)

type UseCases struct {
	UserUseCase usecase.UserUseCaseInterface
}

type UseCasesInit struct {
	dc *DependencyContainer

	useCases      *UseCases
	initFunctions []func(*DependencyContainer) error
}

func NewUseCasesInit(dc *DependencyContainer) *UseCasesInit {
	useCases := &UseCases{}
	return &UseCasesInit{
		dc:       dc,
		useCases: useCases,
		initFunctions: []func(*DependencyContainer) error{
			useCases.initUserUseCase,
		},
	}
}

func (i *UseCasesInit) InitServices() error {
	for _, initFunc := range i.initFunctions {
		if err := initFunc(i.dc); err != nil {
			return err
		}
	}

	i.dc.UseCases = i.useCases
	return nil
}

func (i *UseCasesInit) StartServices() error {
	return nil
}

func (i *UseCases) initUserUseCase(dc *DependencyContainer) error {
	useCase, err := userUseCase.NewBuilder().
		UserRepo(dc.Repositories.UserRepository).
		JwtService(dc.JWTService).
		Build()
	if err != nil {
		dc.LogError(err)
		return err
	}

	i.UserUseCase = useCase
	return nil
}
