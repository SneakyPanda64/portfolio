package ws

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func Upgrader(c *fiber.Ctx) error {
	if websocket.IsWebSocketUpgrade(c) {
		c.Locals("allowed", true)
		c.Locals("ip", c.Get("cf-connecting-ip"))
		c.Locals("countr ", c.Get("cf-ipcountry"))
		return c.Next()
	}
	return fiber.ErrUpgradeRequired
}
