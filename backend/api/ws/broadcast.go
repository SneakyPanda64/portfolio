package ws

import "github.com/sirupsen/logrus"

func BroadcastMessages() {
	for {
		// Grab the next message from the broadcast channel
		message := <-broadcast

		// Send the message to all connected clients
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
