package taskUseCase

import (
	userGatewayInterface "common/gateways/user/interface"
	"infrastructure/errors"
	kafkaClientInterface "infrastructure/kafka/interface"
	repositoryInterface "task-service/boundary/repository"
)

type Builder struct {
	taskUseCase *TaskUseCase
	errors      *errors.Errors
}

func NewBuilder() *Builder {
	return &Builder{
		taskUseCase: &TaskUseCase{},
		errors:      errors.NewErrors(),
	}
}

func (b *Builder) TaskRepo(taskRepo repositoryInterface.TaskRepository) *Builder {
	b.taskUseCase.taskRepo = taskRepo
	return b
}

func (b *Builder) KafkaClient(kafkaClient kafkaClientInterface.KafkaClient) *Builder {
	b.taskUseCase.kafkaClient = kafkaClient
	return b
}

func (b *Builder) UserGateway(userGateway userGatewayInterface.UserGateway) *Builder {
	b.taskUseCase.userGateway = userGateway
	return b
}

func (b *Builder) Build() (*TaskUseCase, error) {
	b.checkRequiredFields()
	if b.errors.IsPresent() {
		return nil, b.errors
	}
	return b.taskUseCase, nil
}

func (b *Builder) checkRequiredFields() {
	if b.taskUseCase.taskRepo == nil {
		b.errors.AddError(errors.NewError("SYS", "TaskUseCaseBuilder: TaskRepository is required"))
	}
	if b.taskUseCase.kafkaClient == nil {
		b.errors.AddError(errors.NewError("SYS", "TaskUseCaseBuilder: KafkaClient is required"))
	}
	if b.taskUseCase.userGateway == nil {
		b.errors.AddError(errors.NewError("SYS", "TaskUseCaseBuilder: UserGateway is required"))
	}
}
