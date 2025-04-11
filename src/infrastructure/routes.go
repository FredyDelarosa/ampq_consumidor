package infrastructure

import (
	"notificaciones/src/application"
	"notificaciones/src/infrastructure/controllers"
	"notificaciones/src/infrastructure/websocket"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, processAlertUseCase *application.ProcessAlertUseCase) {
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	alertController := controllers.NewAlertController(processAlertUseCase.Repo)

	router.GET("/alerts", alertController.GetAllAlerts)
	router.GET("/ws", func(c *gin.Context) {
		websocket.HandleWebSocket(c.Writer, c.Request)
	})
}
