package kafkaGateway

import (
	"context"
)

type MessageGateway interface {
	SendMessage(ctx context.Context, topic, message string) error
	ReadMessage(ctx context.Context, topic string) (string, error)
}
