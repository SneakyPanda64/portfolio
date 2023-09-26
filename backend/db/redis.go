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
	opt, err := redis.ParseURL(os.Getenv("REDIS_URI"))
	if err != nil {
		logrus.Error(err)
		return
	}
	client := redis.NewClient(opt)
	_, err = client.Set(context.Background(), "foo", "bar", 1).Result()
	if err != nil {
		logrus.Error(err)
		return
	}
	logrus.Info("Connected to redis db")
	Redis_client = client
}
