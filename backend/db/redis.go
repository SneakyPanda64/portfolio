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
	res, err := client.Ping(context.Background()).Result()
	if err != nil {
		return err
	}
	logrus.Info(res)
	logrus.Info("Connected to redis db")
	Redis_client = client
	return nil
}
