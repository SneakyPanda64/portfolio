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
		Addr: "na",
		// Username: os.Getenv("REDIS_USERNAME"),
		// Password: os.Getenv("REDIS_PASSWORD"), // no password set
		DB: 0, // use default DB
	})
	logrus.Infof("%s, %s, %s", os.Getenv("REDIS_URI"), os.Getenv("REDIS_USERNAME"), os.Getenv("REDIS_PASSWORD"))
	err := client.Ping(context.Background()).Err()
	if err != nil {
		return err
	}
	logrus.Info("Connected to redis db")
	Redis_client = client
	return nil
}
