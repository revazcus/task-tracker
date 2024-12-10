package mongoInterface

import "context"

type MongoRepository interface {
	Create(ctx context.Context, collectionName string, document interface{}) (string, error)
	GetByID(ctx context.Context, collectionName string, id string, resultModel interface{}) error
	GetAll(ctx context.Context, collectionName string) ([]interface{}, error)
	Update(ctx context.Context, collectionName string, id string, update interface{}) error
	Delete(ctx context.Context, collectionName string, id string) error
}
