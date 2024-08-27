package mongo

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Makcumblch/asynchronous-like-counter/internal/domain/counter"
	"github.com/Makcumblch/asynchronous-like-counter/internal/util/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

func NewMongo(config config.MongoConfig) *Mongo {
	options := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%s", config.User, config.Pass, config.IP, config.Port))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options)
	if err != nil {
		log.Fatal("Ошибка подключения к MongoDB: ", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Ошибка подключения к MongoDB (Ping): ", err)
	}

	collection := client.Database("async-like-counter").Collection("counter")
	res := collection.FindOne(context.Background(), bson.D{})
	if res.Err() != nil {
		_, err := collection.InsertOne(context.Background(), counter.Counter{})
		if err != nil {
			log.Fatal("Ошибка инициализации базы: ", err)
		}
	}

	mongo := &Mongo{
		Client:     client,
		Collection: collection,
	}
	return mongo
}

func (m *Mongo) Close(ctx context.Context) error {
	return m.Client.Disconnect(ctx)
}
