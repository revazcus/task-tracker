package mongoInterface

import (
	"context"
	mongoModel "task-tracker/infrastructure/mongo/model"
)

type MongoRepository interface {
	InsertOne(ctx context.Context, collectionName string, data interface{}) error

	FindOne(ctx context.Context, collectionName string, filter, resultModel interface{}) error

	UpdateOne(ctx context.Context, collectionName string, filter, data interface{}) error

	DeleteOne(ctx context.Context, collectionName string, filter interface{}) error

	CreateIndex(ctx context.Context, index *mongoModel.DBIndex) (string, error)
	CollectionIndexes(ctx context.Context, collection string) (map[string]*mongoModel.DBIndex, error)
	TryCreateIndex(ctx context.Context, index *mongoModel.DBIndex) error
}
