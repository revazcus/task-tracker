package kafkaService

import (
	"context"
	"fmt"
	"github.com/IBM/sarama"
	loggerInterface "github.com/revazcus/task-tracker/backend/infrastructure/logger/interface"
)

type KafkaService struct {
	producer  sarama.AsyncProducer
	successes chan *sarama.ProducerMessage
	errors    chan *sarama.ProducerError
	consumer  sarama.Consumer
	admin     sarama.ClusterAdmin
	logger    loggerInterface.Logger
}

// NewKafkaService Переписать на билдер
func NewKafkaService(brokers []string, logger loggerInterface.Logger) (*KafkaService, error) {
	config := sarama.NewConfig()
	config.Version = sarama.V3_9_0_0
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true

	producer, err := sarama.NewAsyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}

	consumer, err := sarama.NewConsumer(brokers)

	successes := make(chan *sarama.ProducerMessage)
	errors := make(chan *sarama.ProducerError)

	go func() {
		for msg := range successes {
			fmt.Printf("Сообщение успешно отправлено: %s", msg.Value)
		}
	}()

	go func() {
		for err := range errors {
			fmt.Printf("Ошибка при отправке сообщения: %v", err)
		}
	}()

	return &KafkaService{
		producer:  producer,
		successes: successes,
		errors:    errors,
		logger:    logger,
	}, nil
}

func (s *KafkaService) SendMessage(ctx context.Context, topic, message string) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	s.producer.Input() <- msg

	select {
	case <-s.successes:
		s.logger.Info(ctx, fmt.Sprintf("Сообщение в Kafka отправлено успешно"))
		return nil
	case err := <-s.errors:
		s.logger.Error(ctx, fmt.Errorf("ошибка при отправке сообщения в Kafka: %v", err))
		return err
	}

}

func (s *KafkaService) Close() error {
	return s.producer.Close()
}

func (s *KafkaService) ReadMessage(ctx context.Context, topic string) (string, error) {
	//TODO implement me
	panic("implement me")
}
