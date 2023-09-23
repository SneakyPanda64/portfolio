package api

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/sneakypanda64/portfolio/api/ws"
)

func RegisterRoutes(app *fiber.App) {
	app.Use("/ws", ws.Upgrader)
	app.Get("/ws/:id", websocket.New(func(c *websocket.Conn) {
		ws.Handler(c)
	}))
	go ws.BroadcastMessages()
}
