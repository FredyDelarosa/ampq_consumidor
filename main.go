package main

import (
	"log"
	"notificaciones/src/infrastructure"

	"github.com/gin-gonic/gin"
)

func main() {
	deps, err := infrastructure.NewDependencies()
	if err != nil {
		log.Fatal("Error al inicializar dependencias:", err)
	}

	r := gin.Default()

	infrastructure.RegisterRoutes(r, deps.ProcessAlertUseCase)

	r.Run(":8081")
}
