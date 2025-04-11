package services

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type RabbitMQFetcher struct{}

func NewRabbitMQFetcher() *RabbitMQFetcher {
	return &RabbitMQFetcher{}
}

func (s *RabbitMQFetcher) FetchAlerts() ([]map[string]interface{}, error) {
	resp, err := http.Get("http://localhost:9090/alerts")
	if err != nil {
		log.Println("Error obteniendo alertas del consumidor:", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error leyendo respuesta del consumidor:", err)
		return nil, err
	}

	var alerts []map[string]interface{}
	if err := json.Unmarshal(body, &alerts); err != nil {
		log.Println("Error decodificando JSON:", err)
		return nil, err
	}

	return alerts, nil
}
