package ws

import "github.com/sirupsen/logrus"

func BroadcastMessages() {
	for {
		message := <-broadcast
		for client := range clients {
			err := client.Conn.WriteMessage(1, message)
			if err != nil {
				logrus.Error("Error writing message:", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
