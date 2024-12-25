package mongoRepo

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"task-tracker/infrastructure/errors"
	mongoModel "task-tracker/infrastructure/mongo/model"
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

func (r *MongoRepo) InsertOne(ctx context.Context, collectionName string, data interface{}) error {
	coll := r.mongoDB.Collection(collectionName)
	result, err := coll.InsertOne(ctx, data)
	if err != nil {
		mongoErr, isMongoErr := err.(mongo.WriteException)
		if isMongoErr {
			for _, mongoWriteErr := range mongoErr.WriteErrors {
				if r.IsDuplicateError(mongoWriteErr) {
					return errors.NewError("Mongo", mongoWriteErr.Message)
				}
			}
		}
		return err
	}

	if _, ok := result.InsertedID.(bson.ObjectID); !ok {
		return mongo.ErrNilDocument
	}

	return nil
}

func (r *MongoRepo) Find(ctx context.Context, collectionName string, results, find interface{}, opt *options.FindOptionsBuilder) error {
	collection := r.mongoDB.Collection(collectionName)
	cursor, err := collection.Find(ctx, find, opt)
	if err != nil {
		return err
	}
	err = cursor.All(ctx, results)
	if err != nil {
		return err
	}
	return nil
}

func (r *MongoRepo) FindOne(ctx context.Context, collectionName string, filter, resultModel interface{}) error {
	collection := r.mongoDB.Collection(collectionName)
	result := collection.FindOne(ctx, filter)
	if err := result.Err(); err != nil {
		return err
	}
	if err := result.Decode(resultModel); err != nil {
		return err
	}
	return nil
}

func (r *MongoRepo) FindOneAndUpdate(ctx context.Context, collectionName string, resultModel, filter, updateData interface{}, opt *options.FindOneAndUpdateOptionsBuilder) error {
	collection := r.mongoDB.Collection(collectionName)
	result := collection.FindOneAndUpdate(ctx, filter, updateData, opt)
	if err := result.Err(); err != nil {
		return err
	}
	if err := result.Decode(resultModel); err != nil {
		return err
	}
	return nil
}

func (r *MongoRepo) UpdateOne(ctx context.Context, collectionName string, filter, data interface{}, opts ...*options.UpdateOptions) error {
	collection := r.mongoDB.Collection(collectionName)
	if _, err := collection.UpdateOne(ctx, filter, data); err != nil {
		return err
	}
	return nil
}

func (r *MongoRepo) FindAll(ctx context.Context, collection string) ([]interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (r *MongoRepo) DeleteOne(ctx context.Context, collectionName string, filter interface{}) error {
	collection := r.mongoDB.Collection(collectionName)
	if _, err := collection.DeleteOne(ctx, filter); err != nil {
		return err
	}
	return nil
}

func (r *MongoRepo) CreateIndex(ctx context.Context, index *mongoModel.DBIndex) (string, error) {
	// Получаем коллекцию из БД по имени
	collection := r.mongoDB.Collection(index.Collection)

	// Создаём пустой слайс, содержащий bson.E (пара "ключ-значение") для каждого поля индекса
	keysName := make([]bson.E, 0)

	// Проходим по слайсу с именами полей, по которым построим индекс
	for _, key := range index.Keys {
		keysName = append(keysName, bson.E{
			Key:   key,
			Value: int(index.Type),
		})
	}

	// Создаём упорядоченный документ BSON, гарантирующий, что ключи будут отсортированы в том порядке, в котором они были добавлены
	keys := bson.D(keysName)

	// Создаём и наполняем структуру для создания индекса
	indexModel := mongo.IndexModel{}
	indexModel.Keys = keys
	indexModel.Options = options.Index().SetName(index.Name)
	if index.Uniq {
		// Если индекс уникальный, то гарантируем, что значения в индексируемых полях коллекции будут уникальными
		indexModel.Options.SetUnique(true)
	}

	// Создаём настроенный индекс в коллекции
	return collection.Indexes().CreateOne(ctx, indexModel)
}

func (r *MongoRepo) CollectionIndexes(ctx context.Context, collection string) (map[string]*mongoModel.DBIndex, error) {
	result := make(map[string]*mongoModel.DBIndex)

	// Получаем коллекцию из БД по имени
	c := r.mongoDB.Collection(collection)

	// Вытаскиваем все индексы из указанной коллекции
	cur, err := c.Indexes().List(ctx)
	if err != nil {
		return result, err
	}

	// Проходимся по курсору пока в нём есть элементы
	for cur.Next(ctx) {
		// Создаём структуру индекса, в которую будет декодироваться информация о нём
		index := &mongoModel.DBIndex{}
		if err := cur.Decode(&index); err != nil {
			return result, err
		}
		result[index.Name] = index
	}
	return result, nil
}

func (r *MongoRepo) TryCreateIndex(ctx context.Context, index *mongoModel.DBIndex) error {
	indexes, err := r.CollectionIndexes(ctx, index.Collection)
	if err != nil {
		return err
	}

	if r.isIndexExist(index, indexes) {
		return nil
	}

	if _, err = r.CreateIndex(ctx, index); err != nil {
		return err
	}

	return nil
}

func (r *MongoRepo) isIndexExist(index *mongoModel.DBIndex, indexes map[string]*mongoModel.DBIndex) bool {
	_, ok := indexes[index.Name]
	return ok
}

func (r *MongoRepo) IsDuplicateError(mongoWriteError mongo.WriteError) bool {
	return mongoWriteError.Code == 11000
}
