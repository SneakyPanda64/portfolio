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

func Connect() error {
	opt, err := redis.ParseURL(os.Getenv("REDIS_URI"))
	if err != nil {
		return err
	}
	client := redis.NewClient(opt)
	_, err = client.Set(context.Background(), "clicks", 5, 0).Result()
	if err != nil {
		return err
	}
	logrus.Info("Connected to redis db")
	Redis_client = client
	return nil
}
