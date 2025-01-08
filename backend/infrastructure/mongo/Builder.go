package mongoRepo

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
	"infrastructure/errors"
	loggerInterface "infrastructure/logger/interface"
)

type Builder struct {
	mongoDb *mongo.Database
	logger  loggerInterface.Logger

	errors *errors.Errors
}

func NewBuilder() *Builder {
	return &Builder{
		errors: errors.NewErrors(),
	}
}

func (b *Builder) MongoDB(mongoDb *mongo.Database) *Builder {
	b.mongoDb = mongoDb
	return b
}

func (b *Builder) Logger(logger loggerInterface.Logger) *Builder {
	b.logger = logger
	return b
}

func (b *Builder) Build() (*MongoRepo, error) {
	b.checkRequiredFields()
	if b.errors.IsPresent() {
		return nil, b.errors
	}

	return b.createFromBuilder(), nil
}

func (b *Builder) checkRequiredFields() {
	if b.mongoDb == nil {
		b.errors.AddError(ErrMongoDbIsRequired)
	}
	if b.logger == nil {
		b.errors.AddError(ErrLoggerIsRequired)
	}
}

func (b *Builder) createFromBuilder() *MongoRepo {
	return &MongoRepo{mongoDB: b.mongoDb, logger: b.logger}
}
