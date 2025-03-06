package repositories

import (
	"notificaciones/src/domain/entities"
)

type AlertRepository interface {
	Create(alert *entities.Alert) error
	GetAll() ([]entities.Alert, error)
}
