package mongoInterface

import (
	"context"
)

type MongoRepository interface {
	InsertOne(ctx context.Context, collectionName string, data interface{}) error
	FindOne(ctx context.Context, collectionName string, filter, resultModel interface{}) error
	UpdateOne(ctx context.Context, collectionName string, filter, data interface{}) error
	DeleteOne(ctx context.Context, collectionName string, filter interface{}) error
}
