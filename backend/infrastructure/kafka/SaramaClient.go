package saramaClient

import (
	"context"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/revazcus/task-tracker/backend/infrastructure/errors"
	"github.com/revazcus/task-tracker/backend/infrastructure/kafka/event"
	groupHenler "github.com/revazcus/task-tracker/backend/infrastructure/kafka/handler"
	loggerInterface "github.com/revazcus/task-tracker/backend/infrastructure/logger/interface"
	"time"
)

type SaramaClient struct {
	admin     sarama.ClusterAdmin
	producer  sarama.AsyncProducer
	consumer  sarama.ConsumerGroup
	successes chan *sarama.ProducerMessage
	errors    chan *sarama.ProducerError
	logger    loggerInterface.Logger
}

func NewSaramaClient(brokers []string, groupId string, logger loggerInterface.Logger) (*SaramaClient, error) {
	admin, err := sarama.NewClusterAdmin(brokers, sarama.NewConfig())
	if err != nil {
		logger.Error(context.Background(), fmt.Errorf("failed to create cluster admin: %w", err))
		return nil, err
	}

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	config.Consumer.Offsets.AutoCommit.Enable = true
	config.Consumer.Offsets.AutoCommit.Interval = 1 * time.Second
	config.Consumer.Offsets.Retention = 24 * time.Hour

	producer, err := sarama.NewAsyncProducer(brokers, config)
	if err != nil {
		logger.Error(context.Background(), fmt.Errorf("failed to create producer: %w", err))
		return nil, err
	}

	consumer, err := sarama.NewConsumerGroup(brokers, groupId, config)
	if err != nil {
		logger.Error(context.Background(), fmt.Errorf("failed to create consumer: %w", err))
		return nil, err
	}

	successes := make(chan *sarama.ProducerMessage)
	errors := make(chan *sarama.ProducerError)

	saramaClient := &SaramaClient{
		admin:     admin,
		producer:  producer,
		consumer:  consumer,
		successes: successes,
		errors:    errors,
		logger:    logger,
	}

	go func() {
		for msg := range producer.Successes() {
			successes <- msg
		}
	}()

	go func() {
		for err := range producer.Errors() {
			errors <- err
		}
	}()

	return saramaClient, nil
}

func (c *SaramaClient) CreateTopic(ctx context.Context, name string, partitions int32, replicationFactor int16) error {
	topicDetail := &sarama.TopicDetail{
		NumPartitions:     partitions,
		ReplicationFactor: replicationFactor,
	}
	if err := c.admin.CreateTopic(name, topicDetail, false); err != nil {
		c.logger.Error(ctx, fmt.Errorf("failed to create topic: %w", err))
		return err
	}
	c.logger.Info(ctx, fmt.Sprintf("topic %s created", name))
	return nil
}

func (c *SaramaClient) DeleteTopic(ctx context.Context, name string) error {
	if err := c.admin.DeleteTopic(name); err != nil {
		c.logger.Error(ctx, fmt.Errorf("failed to delete topic: %w", err))
		return err
	}
	c.logger.Info(ctx, fmt.Sprintf("topic %s deleted", name))
	return nil
}

func (c *SaramaClient) SendMessage(ctx context.Context, topic string, eventNotification *event.EventNotification) error {
	msgBytes, err := eventNotification.ToBytes()
	if err != nil {
		c.logger.Error(ctx, fmt.Errorf("failed to convert event notification to bytes: %w", err))
	}

	select {
	case c.producer.Input() <- &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(msgBytes)}:
		c.logger.Info(ctx, fmt.Sprintf("message enqueued to topic %s", topic))
	case <-ctx.Done():
		c.logger.Error(ctx, fmt.Errorf("context canceled while sending message to topic %s", topic))
		return ctx.Err()
	}

	select {
	case <-c.successes:
		c.logger.Info(ctx, fmt.Sprintf("message sent to topic %s", topic))
	case err := <-c.errors:
		c.logger.Error(ctx, fmt.Errorf("failed to send message to topic %s: %w", topic, err.Err))
		return err.Err
	case <-ctx.Done():
		c.logger.Error(ctx, fmt.Errorf("context canceled while waiting for message result for topic %s", topic))
		return ctx.Err()
	}

	return nil
}

func (c *SaramaClient) ReadMessage(ctx context.Context, topic string) (*event.EventNotification, error) {

	// Создаём и инициализируем обработчик
	handler := &groupHenler.ConsumerGroupHandler{
		MessageChan: make(chan *sarama.ConsumerMessage), // Канал для получения сообщений
	}

	// Запускаем потребление сообщений в отдельной горутине
	go func() {
		c.logger.Info(ctx, fmt.Sprintf("Starting to consume messages from topic %s", topic))
		for {
			// Пытаемся потреблять сообщения из Kafka
			if err := c.consumer.Consume(ctx, []string{topic}, handler); err != nil {
				c.logger.Error(ctx, fmt.Errorf("failed to consume messages from topic %s: %w", topic, err))
				return
			}

			// Если контекст был отменен, останавливаем потребление
			if ctx.Err() != nil {
				c.logger.Info(ctx, fmt.Sprintf("Context cancelled, stopping consumption of topic %s", topic))
				return
			}
		}
	}()

	// Ждем первое сообщение из канала
	select {
	case msg := <-handler.MessageChan:
		// Если сообщение получено, конвертируем его в EventNotification
		c.logger.Info(ctx, fmt.Sprintf("message read from topic %s", topic))
		var eventNotification event.EventNotification
		if err := eventNotification.FromBytes(msg.Value); err != nil {
			c.logger.Error(ctx, fmt.Errorf("failed to convert message to event notification: %w", err))
			return nil, err
		}

		// Возвращаем прочитанное сообщение
		return &eventNotification, nil
	case <-ctx.Done():
		// В случае отмены контекста
		c.logger.Error(ctx, fmt.Errorf("context canceled while reading message from topic %s", topic))
		return nil, ctx.Err()
	}
}

func (c *SaramaClient) Close() error {
	errs := errors.NewErrors()

	if err := c.admin.Close(); err != nil {
		errs.AddError(errors.NewError("SYS", fmt.Sprintf("failed to close cluster admin: %e", err)))
	}

	if err := c.producer.Close(); err != nil {
		errs.AddError(errors.NewError("SYS", fmt.Sprintf("failed to close producer: %e", err)))
	}

	if err := c.consumer.Close(); err != nil {
		errs.AddError(errors.NewError("SYS", fmt.Sprintf("failed to close consumer: %e", err)))
	}

	if errs.IsPresent() {
		return errs
	}

	c.logger.Info(context.Background(), "sarama client closed successfully")
	return nil
}
