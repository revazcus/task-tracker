package kafkaBroker

import (
	"context"
	"fmt"
	userEventTypes "github.com/revazcus/task-tracker/backend/common/broker/user/event"
	userTopicNames "github.com/revazcus/task-tracker/backend/common/broker/user/topic"
	kafkaClientInterface "github.com/revazcus/task-tracker/backend/infrastructure/kafka/interface"
	loggerInterface "github.com/revazcus/task-tracker/backend/infrastructure/logger/interface"
	"github.com/revazcus/task-tracker/backend/task-service/boundary/domain/usecase"
)

type EventListener struct {
	kafkaClient  kafkaClientInterface.KafkaClient
	taskUserCase usecase.TaskUseCaseInterface
	logger       loggerInterface.Logger
}

func (l *EventListener) Listen(ctx context.Context) {
	for {
		eventNotification, err := l.kafkaClient.ReadMessage(ctx, userTopicNames.InfoTopic)
		if err != nil {
			l.logger.Error(ctx, fmt.Errorf("failed to read message: %w", err))
			continue
		}

		switch eventType.String() {
		case userEventTypes.EmailUpd:

		case userEventTypes.ProfileUpd:

		default:
			l.logger.Error(ctx, fmt.Errorf("unknown event type: %s", eventType.String()))
			continue
		}
	}
}
