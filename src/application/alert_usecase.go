package application

import (
	"encoding/json"
	"fmt"
	"log"
	"notificaciones/src/core"
	"notificaciones/src/domain/entities"
	"notificaciones/src/domain/repositories"
)

type ProcessAlertUseCase struct {
	Repo repositories.AlertRepository
}

func NewProcessAlertUseCase(repo repositories.AlertRepository) *ProcessAlertUseCase {
	return &ProcessAlertUseCase{Repo: repo}
}

func (uc *ProcessAlertUseCase) StartListening() {
	if core.RabbitChannel == nil {
		log.Println("ni para conectarte al rabbit sirves")
		return
	}

	log.Println("esperemos que funcione")

	msgs, err := core.RabbitChannel.Consume(
		"sensor_alerts", "", true, false, false, false, nil,
	)
	if err != nil {
		log.Fatal("no funciono :(", err)
	}

	go func() {
		for msg := range msgs {
			var event map[string]interface{}
			if err := json.Unmarshal(msg.Body, &event); err != nil {
				log.Println("papi ni crear un JSON?", err)
				continue
			}

			message := fmt.Sprintf("Movimiento detectado en zona %s", event["zone"])
			alert := entities.Alert{Message: message}

			if err := uc.Repo.Create(&alert); err != nil {
				log.Println("Error guardando alerta en MySQL:", err)
			} else {
				log.Println("Alerta guardada:", message)
			}
		}
	}()
}
