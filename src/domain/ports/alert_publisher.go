package ports

import "notificaciones/src/domain/entities"

type AlertPublisher interface {
	PublishProcessedAlert(alert *entities.Alert) error
}
