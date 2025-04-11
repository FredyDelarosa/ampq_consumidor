package application

import (
	"fmt"
	"log"
	"notificaciones/src/domain/entities"
	"notificaciones/src/domain/ports"
	"notificaciones/src/domain/repositories"
	"time"
)

type ProcessAlertUseCase struct {
	Repo      repositories.AlertRepository
	Fetcher   ports.AlertFetcher
	Publisher ports.AlertPublisher
	Notifier  ports.AlertNotifier
}

func NewProcessAlertUseCase(
	repo repositories.AlertRepository,
	fetcher ports.AlertFetcher,
	publisher ports.AlertPublisher,
	notifier ports.AlertNotifier,
) *ProcessAlertUseCase {
	return &ProcessAlertUseCase{
		Repo:      repo,
		Fetcher:   fetcher,
		Publisher: publisher,
		Notifier:  notifier,
	}
}

func (uc *ProcessAlertUseCase) StartFetchingAlerts() {
	for {
		time.Sleep(5 * time.Second)

		alerts, err := uc.Fetcher.FetchAlerts()
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

				if err := uc.Notifier.Notify(&alert); err != nil {
					log.Println("Error enviando por WebSocket:", err)
				}
			}

			if err := uc.Publisher.PublishProcessedAlert(&alert); err != nil {
				log.Println("Error publicando alerta procesada:", err)
			} else {
				log.Println("Alerta procesada publicada:", message)
			}
		}
	}
}
