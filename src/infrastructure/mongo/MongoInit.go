package mongoRepo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"time"
)

func InitMongoDatabase(url, mongoDBName string) (*mongo.Database, error) {

	mongoClient, err := mongo.Connect(options.Client().ApplyURI(url))
	if err != nil {
		return nil, err
	}

	pingTimeout := time.Now().Add(10 * time.Second)
	ctx, cancelFunc := context.WithDeadline(context.Background(), pingTimeout)
	defer cancelFunc()

	if err := mongoClient.Ping(ctx, nil); err != nil {
		return nil, err
	}

	mongoDB := mongoClient.Database(mongoDBName)

	fmt.Println("Подключение к MongoDB успешно!")

	return mongoDB, nil
}
