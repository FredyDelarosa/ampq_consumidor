package application

import (
	"fmt"
	"log"
	"notificaciones/src/domain/entities"
	"notificaciones/src/domain/repositories"
	"notificaciones/src/infrastructure/services"
	"time"
)

type ProcessAlertUseCase struct {
	Repo   repositories.AlertRepository
	Rabbit *services.RabbitMQService
}

func NewProcessAlertUseCase(repo repositories.AlertRepository, rabbit *services.RabbitMQService) *ProcessAlertUseCase {
	return &ProcessAlertUseCase{Repo: repo, Rabbit: rabbit}
}

func (uc *ProcessAlertUseCase) StartFetchingAlerts() {
	log.Println("📡 Conectando al consumidor en http://localhost:9090/alerts...")

	for {
		time.Sleep(5 * time.Second)

		alerts, err := uc.Rabbit.FetchAlerts()
		if err != nil {
			continue
		}

		for _, event := range alerts {
			zone, ok := event["zone"].(string)
			if !ok {
				log.Println("Error: El campo 'zone' no es un string válido.")
				continue
			}

			message := fmt.Sprintf("Movimiento detectado en zona %s", zone)
			alert := entities.Alert{Message: message}

			if err := uc.Repo.Create(&alert); err != nil {
				log.Println("Error guardando alerta en MySQL:", err)
			} else {
				log.Println("Alerta guardada:", message)
			}
		}
	}
}
