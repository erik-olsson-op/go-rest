package handlers

import (
	"fmt"
	"github.com/erik-olsson-op/go-rest/internal/logger"
	"github.com/erik-olsson-op/go-rest/internal/models"
	"github.com/erik-olsson-op/go-rest/internal/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// SignUpV1
// @Summary Signup new user
// @Description Signup new user
// @Accept  json
// @Produce  json
// @Success 201
// @Failure 400 {object} models.ErrMessage "failed request"
// @Param   models.Credentials body models.Credentials true "Credentials"
// @Router /api/v1/users [post]
func SignUpV1(ctx *gin.Context) {
	var credentials models.Credentials
	err := ctx.ShouldBindJSON(&credentials)
	if err != nil {
		logger.Logger.Warning(err)
		ctx.JSON(http.StatusBadRequest, models.NewErrMessage(fmt.Sprintf("%v", err)))
		return
	}

	id, err := credentials.Save()
	if err != nil {
		logger.Logger.Warning(err)
		ctx.JSON(http.StatusInternalServerError, models.NewErrMessage("failed to save user to db"))
		return
	}
	ctx.Header("Location", fmt.Sprintf("/api/v1/users/%d", id))
	ctx.Status(http.StatusCreated)
}

// LoginV1
// @Summary Login user (Get JWT Token)
// @Description Login user (Get JWT Token)
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 400 {object} models.ErrMessage "failed request"
// @Param   models.Credentials body models.Credentials true "Email and Password"
// @Router /api/v1/users/login [post]
func LoginV1(ctx *gin.Context) {
	var credentials models.Credentials
	err := ctx.BindJSON(&credentials)
	if err != nil {
		logger.Logger.Warning(err)
		ctx.JSON(http.StatusBadRequest, models.NewErrMessage(fmt.Sprintf("%v", err)))
		return
	}

	userId, email, err := models.ValidateCredentials(&credentials)
	if err != nil {
		logger.Logger.Warning(err)
		ctx.JSON(http.StatusBadRequest, models.NewErrMessage(fmt.Sprintf("%v", err)))
		return
	}

	token, err := util.GenerateToken(email, strconv.FormatInt(userId, 10))
	if err != nil {
		logger.Logger.Warning(err)
		ctx.JSON(http.StatusInternalServerError, models.NewErrMessage("something went wrong"))
		return
	}

	ctx.JSON(200, gin.H{"token": token})
}
