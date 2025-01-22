package initService

import (
	"context"
	saramaClient "infrastructure/kafka"
	kafkaClientInterface "infrastructure/kafka/interface"
)

type KafkaClient struct {
	SaramaClient kafkaClientInterface.KafkaClient
}

type KafkaClientInit struct {
	dc *DependencyContainer
}

func NewKafkaInit(dc *DependencyContainer) *KafkaClientInit {
	return &KafkaClientInit{
		dc: dc,
	}
}

func (i *KafkaClientInit) InitServices() error {
	kafkaBroker, err := i.dc.EnvRegistry.GetEnv(KafkaBroker)
	if err != nil {
		i.dc.LogError(err)
		return err
	}

	groupId, err := i.dc.EnvRegistry.GetEnv(KafkaGroupId)
	if err != nil {
		i.dc.LogError(err)
		return err
	}

	saramaClient, err := saramaClient.NewSaramaClient([]string{kafkaBroker}, groupId, i.dc.Logger)
	if err != nil {
		i.dc.LogError(err)
		return err
	}

	i.dc.KafkaClient = saramaClient

	return nil
}

func (i *KafkaClientInit) StartServices() error {
	topic, err := i.dc.EnvRegistry.GetEnv(KafkaTopic)
	if err != nil {
		i.dc.LogError(err)
		return err
	}

	if err := i.dc.KafkaClient.CreateTopic(context.Background(), topic, 3, 1); err != nil {
		i.dc.Logger.Error(context.Background(), err)
	}

	return nil
}
