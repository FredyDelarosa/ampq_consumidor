package websocket

import (
	"log"

	"github.com/gorilla/websocket"
)

// Use the Clients variable declared in websocket.go

func BroadcastToClients(data []byte) {
	for client := range Clients {
		err := client.WriteMessage(websocket.TextMessage, data)
		if err != nil {
			log.Println("Error enviando a cliente WebSocket:", err)
			client.Close()
			delete(Clients, client)
		}
	}
}
