package mongoInterface

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	mongoModel "infrastructure/mongo/model"
)

type MongoRepository interface {
	InsertOne(ctx context.Context, collectionName string, data interface{}) error

	Find(ctx context.Context, collectionName string, results, find interface{}, opt *options.FindOptionsBuilder) error
	FindOne(ctx context.Context, collectionName string, filter, resultModel interface{}) error
	FindOneAndUpdate(ctx context.Context, collectionName string, resultModel, filter, updateData interface{}, opt *options.FindOneAndUpdateOptionsBuilder) error

	UpdateOne(ctx context.Context, collectionName string, filter, data interface{}, opts ...*options.UpdateOneOptions) error

	DeleteOne(ctx context.Context, collectionName string, filter interface{}) error

	CreateIndex(ctx context.Context, index *mongoModel.DBIndex) (string, error)
	CollectionIndexes(ctx context.Context, collection string) (map[string]*mongoModel.DBIndex, error)
	TryCreateIndex(ctx context.Context, index *mongoModel.DBIndex) error
}
