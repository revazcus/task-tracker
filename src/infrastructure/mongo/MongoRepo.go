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

func (m *MongoRepo) Create(ctx context.Context, collectionName string, document interface{}) (string, error) {
	coll := m.mongoDB.Collection(collectionName)
	result, err := coll.InsertOne(ctx, document)
	if err != nil {
		return "", err
	}

	// TODO видится, чтобы сразу приводить к string нужно сгенерировать свои id, а не полагаться на монговское _id
	id, ok := result.InsertedID.(bson.ObjectID)
	if !ok {
		return "", mongo.ErrNilDocument
	}
	return id.Hex(), nil
}

func (m *MongoRepo) GetByID(ctx context.Context, collectionName string, id string, resultModel interface{}) error {
	collection := m.mongoDB.Collection(collectionName)
	filter := bson.M{"_id": id}
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

func (m *MongoRepo) GetAll(ctx context.Context, collection string) ([]interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MongoRepo) Update(ctx context.Context, collection string, id string, update interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (m *MongoRepo) Delete(ctx context.Context, collection string, id string) error {
	//TODO implement me
	panic("implement me")
}
