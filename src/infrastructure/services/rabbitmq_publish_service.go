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

type RabbitMQPublisher struct{}

func NewRabbitMQPublisher() *RabbitMQPublisher {
	return &RabbitMQPublisher{}
}

func (s *RabbitMQPublisher) PublishProcessedAlert(alert *entities.Alert) error {
	if core.RabbitChannel == nil {
		log.Println("no se conecto a rabbit")
		return nil
	}

	_, err := core.RabbitChannel.QueueDeclare(
		"processed_alerts",
		true, false, false, false, nil,
	)
	if err != nil {
		log.Println("error al declarar la cola processed_alerts:", err)
		return err
	}

	body, _ := json.Marshal(alert)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = core.RabbitChannel.PublishWithContext(ctx,
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
