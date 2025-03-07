package services

import (
	"context"
	"encoding/json"
	"log"
	"notificaciones/src/core"
	"notificaciones/src/domain/entities"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

type RabbitMQPublishService struct{}

func NewRabbitMQPublishService() *RabbitMQPublishService {
	return &RabbitMQPublishService{}
}

func (s *RabbitMQPublishService) PublishProcessedAlert(alert *entities.Alert) error {
	if core.RabbitChannel == nil {
		log.Println("no se conecto a rabbit")
		return nil
	}

	body, _ := json.Marshal(alert)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := core.RabbitChannel.PublishWithContext(ctx,
		"", "processed_alerts", false, false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)

	if err != nil {
		log.Println("error al enviar la alerta", err)
	} else {
		log.Println("se envio la alerta", string(body))
	}

	return err
}
