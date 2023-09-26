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
		Password: "", //
		// Username: os.Getenv("REDIS_USERNAME"),
		// Password: os.Getenv("REDIS_PASSWORD"), // no password set
		DB: 0, // use default DB
	})
	logrus.Infof("%s, %s, %s", os.Getenv("REDIS_URI"), os.Getenv("REDIS_USERNAME"), os.Getenv("REDIS_PASSWORD"))
	res, err := client.Ping(context.Background()).Result()
	if err != nil {
		return err
	}
	logrus.Info(res)
	res2, err := client.Info(context.Background()).Result()
	if err != nil {
		return err
	}
	logrus.Info(res2)
	logrus.Info("Connected to redis db")
	Redis_client = client
	return nil
}
