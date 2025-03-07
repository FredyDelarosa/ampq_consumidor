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
	Repo           repositories.AlertRepository
	Rabbit         *services.RabbitMQService
	PublishService *services.RabbitMQPublishService
}

func NewProcessAlertUseCase(repo repositories.AlertRepository, rabbit *services.RabbitMQService, publishService *services.RabbitMQPublishService) *ProcessAlertUseCase {
	return &ProcessAlertUseCase{
		Repo:           repo,
		Rabbit:         rabbit,
		PublishService: publishService,
	}
}

func (uc *ProcessAlertUseCase) StartFetchingAlerts() {
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

			err = uc.PublishService.PublishProcessedAlert(&alert)
			if err != nil {
				log.Println("Error publicando alerta procesada:", err)
			} else {
				log.Println("Alerta procesada publicada:", message)
			}
		}
	}
}
