package mongoMock

import (
	"context"
	mongoModel "github.com/revazcus/task-tracker/backend/infrastructure/mongo/model"
	commonMock "github.com/revazcus/task-tracker/backend/infrastructure/testing/mock"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoRepoMock struct {
	*commonMock.BaseMock
}

func NewMongoRepoMock() *MongoRepoMock {
	return &MongoRepoMock{BaseMock: commonMock.NewBaseMock()}
}

func (m *MongoRepoMock) InsertOne(ctx context.Context, collectionName string, data interface{}) error {
	_, err := m.ProcessMethod("Save", collectionName, data)
	return err
}

func (m *MongoRepoMock) Find(ctx context.Context, collectionName string, results, find interface{}, opt *options.FindOptionsBuilder) error {
	_, err := m.ProcessMethod("Find", collectionName, results, find, opt)
	return err
}

func (m *MongoRepoMock) FindOne(ctx context.Context, collectionName string, filter, resultModel interface{}) error {
	_, err := m.ProcessMethod("FindOne", collectionName, filter, resultModel)
	return err
}

func (m *MongoRepoMock) FindOneAndUpdate(ctx context.Context, collectionName string, resultModel, filter, updateData interface{}, opt *options.FindOneAndUpdateOptionsBuilder) error {
	_, err := m.ProcessMethod("FindOneAndUpdate", collectionName, filter, resultModel)
	return err
}

func (m *MongoRepoMock) UpdateOne(ctx context.Context, collectionName string, filter, data interface{}, opts ...*options.UpdateOptions) error {
	//TODO implement me
	panic("implement me")
}

func (m *MongoRepoMock) DeleteOne(ctx context.Context, collectionName string, filter interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (m *MongoRepoMock) CreateIndex(ctx context.Context, index *mongoModel.DBIndex) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MongoRepoMock) CollectionIndexes(ctx context.Context, collection string) (map[string]*mongoModel.DBIndex, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MongoRepoMock) TryCreateIndex(ctx context.Context, index *mongoModel.DBIndex) error {
	//TODO implement me
	panic("implement me")
}
