package kafkaBroker

import (
	userEventTypes "common/broker/user/event"
	userTopicNames "common/broker/user/topic"
	"context"
	"fmt"
	kafkaClientInterface "infrastructure/kafka/interface"
	loggerInterface "infrastructure/logger/interface"
	"task-service/src/boundary/domain/usecase"
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
