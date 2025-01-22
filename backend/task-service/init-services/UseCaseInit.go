package initService

import (
	"task-service/boundary/domain/usecase"
	taskUseCase "task-service/domain/usecase/task"
)

type UseCases struct {
	TaskUseCase usecase.TaskUseCaseInterface
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
			useCases.initTaskUseCase,
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

func (i *UseCases) initTaskUseCase(dc *DependencyContainer) error {
	useCase, err := taskUseCase.NewBuilder().
		TaskRepo(dc.Repositories.TaskRepository).
		KafkaClient(dc.KafkaClient).
		UserGateway(dc.GrpcGateways.UserGrpcGateway).
		Build()
	if err != nil {
		dc.LogError(err)
		return err
	}

	i.TaskUseCase = useCase
	return nil
}
