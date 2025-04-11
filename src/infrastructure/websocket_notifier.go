package infrastructure

import (
	"encoding/json"
	"notificaciones/src/domain/entities"
	"notificaciones/src/domain/ports"
	"notificaciones/src/infrastructure/websocket"
)

type WebSocketNotifier struct{}

func NewWebSocketNotifier() ports.AlertNotifier {
	return &WebSocketNotifier{}
}

func (wsn *WebSocketNotifier) Notify(alert *entities.Alert) error {
	message, err := json.Marshal(alert)
	if err != nil {
		return err
	}
	websocket.BroadcastToClients(message)
	return nil
}
