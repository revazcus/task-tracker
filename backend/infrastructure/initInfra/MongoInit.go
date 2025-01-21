package initInfra

import (
	mongoRepo "infrastructure/mongo"
)

type MongoInit struct {
}

func NewMongoInit() *MongoInit {
	return &MongoInit{}
}

func (i *MongoInit) InitInfra(ic *InfraContainer) error {
	url, err := ic.EnvRegistry.GetEnv(MongoURIEnv)
	if err != nil {
		ic.LogError(err)
		return err
	}

	dbName, err := ic.EnvRegistry.GetEnv(MongoDBNameEnv)
	if err != nil {
		ic.LogError(err)
		return err
	}

	mongoDB, err := mongoRepo.InitMongoDatabase(url, dbName)
	if err != nil {
		ic.LogError(err)
		return err
	}

	mongoRepository, err := mongoRepo.NewBuilder().
		MongoDB(mongoDB).
		Logger(ic.Logger).
		Build()
	if err != nil {
		ic.LogError(err)
		return err
	}

	ic.MongoRepo = mongoRepository
	return nil
}

func (i *MongoInit) StartInfra(ic *InfraContainer) error {
	return nil
}
