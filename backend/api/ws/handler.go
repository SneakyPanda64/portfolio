package ws

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/gofiber/contrib/websocket"
	"github.com/sirupsen/logrus"
	"github.com/sneakypanda64/portfolio/models"
	"github.com/sneakypanda64/portfolio/src/actions"
	"github.com/sneakypanda64/portfolio/src/heatmap"
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
	for {
		if _, msg, err = c.ReadMessage(); err != nil {
			logrus.Error(err)
			break
		}
		var m *models.Response
		log.Printf("recv: %s", msg)
		json.Unmarshal(msg, &m)
		if m.Action == "click" {
			logrus.Info(c.Locals("ip"), c.Locals("country"))
			ip := strings.ReplaceAll(fmt.Sprintf("%s", c.Locals("ip")), " ", "-")
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
			hm, err := heatmap.GetHeatmap()
			if err != nil {
				logrus.Error(err)
			}
			body, err := json.Marshal(hm)
			if err != nil {
				logrus.Error(err)
			}
			resp := &models.SendData{
				Type:  "stats",
				Value: string(body),
			}
			fbody, err := json.Marshal(resp)
			if err != nil {
				logrus.Error(err)
			}
			c.WriteMessage(1, fbody)

			// c.WriteMessage(1, )
		}
		logrus.Infof("Recieved: action: %s, value: %s", m.Action, m.Value)
	}
	clients[c] = false
	return nil
}
