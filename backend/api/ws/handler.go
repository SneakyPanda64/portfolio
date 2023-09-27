package ws

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/gofiber/contrib/websocket"
	"github.com/sirupsen/logrus"
	"github.com/sneakypanda64/portfolio/models"
	"github.com/sneakypanda64/portfolio/src/actions"
	"github.com/sneakypanda64/portfolio/src/heatmap"
	"github.com/sneakypanda64/portfolio/src/limiting"
)

var (
	clients   = make(map[*websocket.Conn]bool)
	broadcast = make(chan []byte)
)

func Handler(c *websocket.Conn) error {
	var (
		// mt  int
		msg []byte
		err error
	)
	clients[c] = true
	hm_body, err := heatmap.GetStats()
	if err == nil {
		c.WriteMessage(1, hm_body)
	}
	for {
		if _, msg, err = c.ReadMessage(); err != nil {
			logrus.Error(err)
			break
		}
		var m *models.Response
		json.Unmarshal(msg, &m)
		ip := strings.ReplaceAll(fmt.Sprintf("%s", c.Locals("ip")), " ", "-")
		ratelimited, err := limiting.RateLimit(fmt.Sprintf("ratelimit:%s", ip), 60)
		if err != nil {
			logrus.Error(err)
			continue
		}
		if ratelimited {
			logrus.Error(errors.New("rate limited"))
			body, err := json.Marshal(models.SendData{
				Type:  "error",
				Value: "rate limited",
			})
			if err != nil {
				logrus.Error(err)
			}
			c.WriteMessage(1, body)
			continue
		}
		if m.Action == "click" {
			clicks, err := actions.Click(fmt.Sprintf("%s", c.Locals("country")), ip)
			if err != nil {
				logrus.Error(err)
				break
			}
			body, err := json.Marshal(models.SendData{
				Type:  "clicks",
				Value: fmt.Sprintf("%d", clicks),
			})
			if err != nil {
				logrus.Error(err)
				break
			}
			broadcast <- body

		}
		if m.Action == "stats" {
			body, err := heatmap.GetStats()
			if err != nil {
				logrus.Error(err)
				break
			}
			c.WriteMessage(1, body)
		}
		logrus.Infof("Action: %s, Value: %s, IP: %s, Country: %s", m.Action, m.Value, ip, c.Locals("country"))
	}
	clients[c] = false
	return nil
}
