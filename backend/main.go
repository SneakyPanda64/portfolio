package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/storage/redis/v3"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/sneakypanda64/portfolio/api"
	"github.com/sneakypanda64/portfolio/db"
	"github.com/sneakypanda64/portfolio/middleware/logger"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		logrus.Error(err)
	}
	app := fiber.New()

	app.Use(limiter.New())

	err = db.Connect()
	if err != nil {
		logrus.Fatal(err)
	}

	// Or extend your config for customization
	port, err := strconv.ParseInt(strings.Split(os.Getenv("REDIS_URI"), ":")[1], 10, 64)
	if err != nil {
		logrus.Fatal(err)
	}
	store := redis.New(redis.Config{
		Host:      strings.Split(os.Getenv("REDIS_URI"), ":")[0],
		Port:      int(port),
		Password:  os.Getenv("REDIS_PASSWORD"),
		Database:  0,
		Reset:     false,
		TLSConfig: nil,
		PoolSize:  10 * runtime.GOMAXPROCS(0),
	})
	err = store.Conn().Ping(context.Background()).Err()
	if err != nil {
		logrus.Fatal(err)
	}
	app.Use(logger.New(logger.Config{}))
	app.Use(limiter.New(limiter.Config{
		Max:        10,
		Expiration: 60 * time.Second,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.Get("cf-connecting-ip")
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.SendStatus(http.StatusTooManyRequests)
		},
		Storage: store,
	}))

	api.RegisterRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
