package infrastructure

import (
	"log"
	"notificaciones/src/application"
	"notificaciones/src/core"
	"notificaciones/src/infrastructure/services"
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

	rabbitService := services.NewRabbitMQService()
	rabbitPublishService := services.NewRabbitMQPublishService()
	mysqlRepo := NewMySQLAlertRepository(db)
	notifier := NewWebSocketNotifier()

	processAlertUseCase := application.NewProcessAlertUseCase(mysqlRepo, rabbitService, rabbitPublishService, notifier)

	go processAlertUseCase.StartFetchingAlerts()

	return &Dependencies{
		ProcessAlertUseCase: processAlertUseCase,
	}, nil
}
