package counter

import (
	"context"

	"github.com/Makcumblch/asynchronous-like-counter/internal/domain/counter"
	"github.com/Makcumblch/asynchronous-like-counter/internal/infra/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

type CounterRepository struct {
	mongo *mongo.Mongo
}

func NewCounterRepository(mongo *mongo.Mongo) *CounterRepository {
	return &CounterRepository{
		mongo: mongo,
	}
}

func (r *CounterRepository) Get() (uint, error) {
	var counter counter.Counter
	cur := r.mongo.Collection.FindOne(context.Background(), bson.D{})
	err := cur.Decode(&counter)
	if err != nil {
		return 0, err
	}
	return counter.Likes, nil
}

func (r *CounterRepository) Increment() error {
	update := bson.D{
		{Key: "$inc", Value: bson.D{
			{Key: "likes", Value: 1},
		}},
	}
	r.mongo.Collection.UpdateOne(context.Background(), bson.D{}, update)
	return nil
}
