package brokers

import (
	"context"
	"fmt"
	kafkaClientInterface "github.com/revazcus/task-tracker/backend/infrastructure/kafka/interface"
	loggerInterface "github.com/revazcus/task-tracker/backend/infrastructure/logger/interface"
	"github.com/revazcus/task-tracker/backend/user-service/boundary/domain/usecase"
)

type EventListener struct {
	kafkaClient kafkaClientInterface.KafkaClient
	userUseCase usecase.UserUseCaseInterface
	logger      loggerInterface.Logger
}

// NewEventListener TODO переписать на билдер
func NewEventListener(kafkaClient kafkaClientInterface.KafkaClient, userUseCase usecase.UserUseCaseInterface, logger loggerInterface.Logger) *EventListener {
	return &EventListener{
		kafkaClient: kafkaClient,
		userUseCase: userUseCase,
		logger:      logger,
	}
}

func (l *EventListener) Listen(ctx context.Context) {
	for {
		eventNotification, err := l.kafkaClient.ReadMessage(ctx, "user-info") // TODO вынести в константу
		if err != nil {
			l.logger.Error(ctx, fmt.Errorf("failed to read message: %w", err))
			continue
		}

		switch *eventNotification.EventType {
		case "TaskCreated":
			user, err := l.userUseCase.GetUserById(ctx, eventNotification.Payload["userId"].(string))
			if err != nil {
				l.logger.Error(ctx, fmt.Errorf("failed to get user by id: %w", err))
				continue
			}
			l.logger.Info(ctx, fmt.Sprintf("User: %v", user))
		}
	}
}
