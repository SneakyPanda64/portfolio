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
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URI"),
		Username: os.Getenv("REDUS_USERNAME"),
		Password: os.Getenv("REDUS_USERNAME"), // no password set
		DB:       0,                           // use default DB
	})
	err := client.Ping(context.Background()).Err()
	if err != nil {
		return err
	}
	err = client.Get(context.Background(), "test").Err()
	if err != nil {
		return err
	}
	logrus.Info("Connected to redis db")
	Redis_client = client
	return nil
}
