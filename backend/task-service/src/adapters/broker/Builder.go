package kafkaBroker

import (
	"infrastructure/errors"
	kafkaClientInterface "infrastructure/kafka/interface"
	loggerInterface "infrastructure/logger/interface"
	"task-service/src/boundary/domain/usecase"
)

type Builder struct {
	eventListener *EventListener
	errors        *errors.Errors
}

func NewBuilder() *Builder {
	return &Builder{
		eventListener: &EventListener{},
		errors:        errors.NewErrors(),
	}
}

func (b *Builder) KafkaClient(kafkaClient kafkaClientInterface.KafkaClient) *Builder {
	b.eventListener.kafkaClient = kafkaClient
	return b
}

func (b *Builder) TaskUseCase(taskUserCase usecase.TaskUseCaseInterface) *Builder {
	b.eventListener.taskUserCase = taskUserCase
	return b
}

func (b *Builder) Logger(logger loggerInterface.Logger) *Builder {
	b.eventListener.logger = logger
	return b
}

func (b *Builder) Build() (*EventListener, error) {
	b.checkRequiredFields()
	if b.errors.IsPresent() {
		return nil, b.errors
	}
	return b.eventListener, nil
}

func (b *Builder) checkRequiredFields() {
	if b.eventListener.kafkaClient == nil {
		b.errors.AddError(errors.NewError("SYS", "EventListenerBuilder: KafkaClient is required"))
	}
	if b.eventListener.taskUserCase == nil {
		b.errors.AddError(errors.NewError("SYS", "EventListenerBuilder: TaskUseCase is required"))
	}
	if b.eventListener.logger == nil {
		b.errors.AddError(errors.NewError("SYS", "EventListenerBuilder: Logger is required"))
	}
}
