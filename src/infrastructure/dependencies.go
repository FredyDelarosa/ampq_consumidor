package infrastructure

import (
	"log"
	"notificaciones/src/application"
	"notificaciones/src/core"
)

type Dependencies struct {
	ProcessAlertUseCase *application.ProcessAlertUseCase
}

func NewDependencies() (*Dependencies, error) {
	db, err := core.InitDB()
	if err != nil {
		return nil, err
	}

	err = core.InitRabbitMQ()
	if err != nil {
		log.Fatal("el rabbit no jala", err)
	}

	mysqlRepo := NewMySQLAlertRepository(db)
	processAlertUseCase := application.NewProcessAlertUseCase(mysqlRepo)

	go processAlertUseCase.StartListening()

	return &Dependencies{
		ProcessAlertUseCase: processAlertUseCase,
	}, nil
}
