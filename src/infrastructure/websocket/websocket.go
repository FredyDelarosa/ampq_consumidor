package websocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var Clients = make(map[*websocket.Conn]bool)

var Upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error actualizando a websocket:", err)
		return
	}
	defer conn.Close()

	Clients[conn] = true
	log.Println("Cliente WebSocket conectado")

	for {
		if _, _, err := conn.NextReader(); err != nil {
			log.Println("Cliente desconectado")
			delete(Clients, conn)
			break
		}
	}
}
