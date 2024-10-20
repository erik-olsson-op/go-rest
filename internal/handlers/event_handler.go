package handlers

import (
	"fmt"
	"github.com/erik-olsson-op/go-rest/internal/logger"
	"github.com/erik-olsson-op/go-rest/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetAllEventsV1
// @Summary Get all events
// @Description Get all events
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Event
// @Failure 500 {object} models.ErrMessage "failed to fetch all events"
// @Router /api/v1/events [get]
func GetAllEventsV1(ctx *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		logger.Logger.Error(err)
		ctx.JSON(http.StatusInternalServerError, models.NewErrMessage("failed to fetch all events"))
		return
	}
	ctx.JSON(http.StatusOK, events)
}

// GetEventByIdV1
// @Summary Get event by id
// @Description Get event by id
// @Accept  json
// @Produce  json
// @Param   id     path    string     true        "Event id"
// @Success 200 {object} models.Event
// @Failure 404 {object} models.ErrMessage "not found - {id}"
// @Router /api/v1/events/{id} [get]
func GetEventByIdV1(ctx *gin.Context) {
	id := ctx.Param("id")
	event, err := models.GetEventById(id)
	if err != nil {
		logger.Logger.Warning(err)
		ctx.JSON(http.StatusNotFound, models.NewErrMessage(fmt.Sprintf("id not found - %s", id)))
		return
	}
	ctx.JSON(http.StatusOK, event)
}

// DeleteEventByIdV1
// @Summary Delete event by id
// @Description Delete event by id
// @Security JWT
// @Accept  json
// @Produce  json
// @Param   id     path    string     true        "Event id"
// @Success 204
// @Failure 404 {object} models.ErrMessage "not found - {id}"
// @Router /api/v1/events/{id} [delete]
func DeleteEventByIdV1(ctx *gin.Context) {
	id := ctx.Param("id")
	err := models.DeleteEventById(id)
	if err != nil {
		logger.Logger.Warning(err)
		ctx.JSON(http.StatusNotFound, models.NewErrMessage(fmt.Sprintf("id not found - %s", id)))
		return
	}
	ctx.Status(http.StatusNoContent)
}

// CreateEventV1
// @Summary Create new event
// @Description Create new event
// @Security JWT
// @Accept  json
// @Produce  json
// @Success 201
// @Failure 400 {object} models.ErrMessage "failed request"
// @Param   models.Event body models.Event true "Event"
// @Router /api/v1/events [post]
func CreateEventV1(ctx *gin.Context) {
	var event = models.Event{}
	err := ctx.ShouldBindJSON(&event)
	if err != nil {
		logger.Logger.Warning(err)
		ctx.JSON(http.StatusBadRequest, models.NewErrMessage("failed to parse JSON to event"))
		return
	}
	event.UserId, _ = strconv.ParseInt(ctx.GetString("userId"), 10, 64)
	id, err := event.Save()
	if err != nil {
		logger.Logger.Error(err)
		ctx.JSON(http.StatusInternalServerError, models.NewErrMessage("failed to save event to db"))
		return
	}
	ctx.Header("Location", fmt.Sprintf("/api/v1/events/%d", id))
	ctx.Status(http.StatusCreated)
}

// UpdateEventByIdV1
// @Summary Update event by id
// @Description Update event by id
// @Security JWT
// @Accept  json
// @Produce  json
// @Param   id     path    string     true        "Event id"
// @Success 204
// @Failure 404 {object} models.ErrMessage "not found - {id}"
// @Router /api/v1/events/{id} [put]
func UpdateEventByIdV1(ctx *gin.Context) {
	id := ctx.Param("id")
	var event = models.Event{}
	err := ctx.ShouldBindJSON(&event)
	if err != nil {
		logger.Logger.Warning(err)
		ctx.JSON(http.StatusBadRequest, models.NewErrMessage("failed to parse JSON to event"))
		return
	}
	err = models.UpdateEventById(id, event)
	if err != nil {
		logger.Logger.Warning(err)
		ctx.JSON(http.StatusNotFound, models.NewErrMessage(fmt.Sprintf("not found - %s", id)))
		return
	}
	ctx.Status(http.StatusNoContent)
}
