package initService

import "infrastructure/envRegistry"

const (
	MongoCollectionEnv envRegistry.EnvKey = "MONGO_COLLECTION"

	KafkaBroker  envRegistry.EnvKey = "KAFKA_BROKER"
	KafkaGroupId envRegistry.EnvKey = "KAFKA_GROUP_ID"
	KafkaTopic   envRegistry.EnvKey = "KAFKA_TOPIC"

	GrpcPort envRegistry.EnvKey = "GRPC_PORT"
)

var initEnvironments = map[envRegistry.EnvKey]bool{
	MongoCollectionEnv: true,
	KafkaBroker:        true,
	KafkaGroupId:       true,
	KafkaTopic:         true,
	GrpcPort:           true,
}

type EnvironmentsInit struct {
	dc *DependencyContainer
}

func NewEnvironmentsInit(dc *DependencyContainer) *EnvironmentsInit {
	return &EnvironmentsInit{
		dc: dc,
	}
}

func (i *EnvironmentsInit) InitServices() error {
	for env := range initEnvironments {
		if err := i.dc.EnvRegistry.FindAndSetEnv(env); err != nil {
			i.dc.LogError(err)
			return err
		}
	}
	return nil
}

func (i *EnvironmentsInit) StartServices() error {
	return nil
}
