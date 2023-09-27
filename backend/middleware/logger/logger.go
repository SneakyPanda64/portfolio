package logger

import (
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Filter func(c *fiber.Ctx) bool
}

func New(config Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()
		if err != nil {
			return err
		}
		stop := time.Now()
		queries := ""
		index := 0
		for name, query := range c.Queries() {
			if index == 0 {
				queries = fmt.Sprintf("?%s=%s", name, query)
			} else {
				queries = fmt.Sprintf("%s&%s=%s", queries, name, query)
			}
			index += 1
		}
		full_path := fmt.Sprintf("%s%s", strings.Trim(string(c.Context().Path()), "/"), queries)
		logrus.WithFields(logrus.Fields{
			"remote":  c.Get("cf-connecting-ip"),
			"host":    string(c.Context().Host()),
			"method":  c.Method(),
			"code":    c.Response().StatusCode(),
			"latency": fmt.Sprintf("%dms", stop.Sub(start).Milliseconds()),
			"country": c.Get("cf-ipcountry"),
		}).Info(full_path)
		return nil
	}
}
