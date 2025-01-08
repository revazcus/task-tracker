package kafkaClientInterface

import (
	"context"
	kafkaEvent "infrastructure/kafka/event"
)

type KafkaClient interface {
	CreateTopic(ctx context.Context, name string, partitions int32, replicationFactor int16) error
	SendMessage(ctx context.Context, topic string, eventNotification *kafkaEvent.EventNotification) error
	ReadMessage(ctx context.Context, topic string) (*kafkaEvent.EventNotification, error)
	DeleteTopic(ctx context.Context, name string) error
	Close() error
}
