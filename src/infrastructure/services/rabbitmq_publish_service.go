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

	// Verificar y crear la cola si no existe
	_, err := core.RabbitChannel.QueueDeclare(
		"processed_alerts", // nombre de la cola
		true,               // durable
		false,              // autoDelete
		false,              // exclusive
		false,              // noWait
		nil,                // arguments
	)
	if err != nil {
		log.Println("error al declarar la cola processed_alerts:", err)
		return err
	}

	body, _ := json.Marshal(alert)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = core.RabbitChannel.PublishWithContext(ctx,
		"",                 // exchange
		"processed_alerts", // routing key (nombre de la cola)
		false,              // mandatory
		false,              // immediate
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
