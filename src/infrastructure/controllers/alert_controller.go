package controllers

import (
	"net/http"
	"notificaciones/src/domain/repositories"

	"github.com/gin-gonic/gin"
)

type AlertController struct {
	repo repositories.AlertRepository
}

func NewAlertController(repo repositories.AlertRepository) *AlertController {
	return &AlertController{repo: repo}
}

func (ctrl *AlertController) GetAllAlerts(c *gin.Context) {
	alerts, err := ctrl.repo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error obteniendo alertas"})
		return
	}

	c.JSON(http.StatusOK, alerts)
}
