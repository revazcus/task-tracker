package mongoRepo

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// MongoRepo Обёртка над монгой
type MongoRepo struct {
	mongoDB *mongo.Database
}

// NewMongoRepo TODO переписать на билдер
func NewMongoRepo(database *mongo.Database) *MongoRepo {
	return &MongoRepo{
		mongoDB: database,
	}
}

func (m *MongoRepo) InsertOne(ctx context.Context, collectionName string, document interface{}) error {
	coll := m.mongoDB.Collection(collectionName)
	result, err := coll.InsertOne(ctx, document)
	if err != nil {
		return err
	}

	if _, ok := result.InsertedID.(bson.ObjectID); !ok {
		return mongo.ErrNilDocument
	}

	return nil
}

func (m *MongoRepo) UpdateOne(ctx context.Context, collectionName string, filter, data interface{}) error {
	collection := m.mongoDB.Collection(collectionName)

	if _, err := collection.UpdateOne(ctx, filter, data); err != nil {
		return err
	}

	return nil
}

func (m *MongoRepo) FindOne(ctx context.Context, collectionName string, filter, resultModel interface{}) error {
	collection := m.mongoDB.Collection(collectionName)

	result := collection.FindOne(ctx, filter)
	err := result.Err()
	if err != nil {
		return err
	}

	err = result.Decode(resultModel)
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoRepo) FindAll(ctx context.Context, collection string) ([]interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MongoRepo) DeleteOne(ctx context.Context, collectionName string, filter interface{}) error {
	collection := m.mongoDB.Collection(collectionName)

	if _, err := collection.DeleteOne(ctx, filter); err != nil {
		return err
	}

	return nil
}
