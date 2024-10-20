package handlers

import (
	"github.com/erik-olsson-op/go-rest/internal/models"
	"github.com/gin-gonic/gin"
)

// Ping
// @Summary Healthcheck request
// @Description do healthcheck
// @Accept  json
// @Produce  json
// @Success 200 {object} models.HealthCheck
// @Router /api/ping [get]
func Ping(ctx *gin.Context) {
	ctx.JSON(200, models.NewHealthCheck("PONG"))
}
