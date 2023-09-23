package db

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

var (
	Redis_client *redis.Client
)

func Connect() {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URI"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := client.Set(context.Background(), "foo", "bar", 0).Result()
	if err != nil {
		logrus.Error(err)
		return
	}
	logrus.Info("Connected to redis db")
	Redis_client = client
}
