package ports

import "notificaciones/src/domain/entities"

type AlertNotifier interface {
	Notify(alert *entities.Alert) error
}
