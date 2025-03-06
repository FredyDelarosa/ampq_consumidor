package infrastructure

import (
	"database/sql"
	"notificaciones/src/domain/entities"
	"notificaciones/src/domain/repositories"
)

type MySQLAlertRepository struct {
	db *sql.DB
}

func NewMySQLAlertRepository(db *sql.DB) repositories.AlertRepository {
	return &MySQLAlertRepository{db: db}
}

func (repo *MySQLAlertRepository) Create(alert *entities.Alert) error {
	query := "INSERT INTO alerts (message) VALUES (?)"
	_, err := repo.db.Exec(query, alert.Message)
	return err
}

func (repo *MySQLAlertRepository) GetAll() ([]entities.Alert, error) {
	query := "SELECT  id, message, created_at FROM alerts ORDER BY created_at DESC"
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var alerts []entities.Alert
	for rows.Next() {
		var alert entities.Alert
		if err := rows.Scan(&alert.ID, &alert.Message, &alert.CreatedAt); err != nil {
			return nil, err
		}
		alerts = append(alerts, alert)
	}
	return alerts, nil
}
