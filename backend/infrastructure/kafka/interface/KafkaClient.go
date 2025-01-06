package kafkaClientInterface

import (
	"context"
	"github.com/revazcus/task-tracker/backend/infrastructure/kafka/event"
)

type KafkaClient interface {
	CreateTopic(ctx context.Context, name string, partitions int32, replicationFactor int16) error
	SendMessage(ctx context.Context, topic string, eventNotification *event.EventNotification) error
	ReadMessage(ctx context.Context, topic string) (*event.EventNotification, error)
	DeleteTopic(ctx context.Context, name string) error
	Close() error
}
