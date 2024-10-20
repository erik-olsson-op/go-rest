package middleware

import (
	"fmt"
	"github.com/erik-olsson-op/go-rest/internal/logger"
	"github.com/erik-olsson-op/go-rest/internal/models"
	"github.com/erik-olsson-op/go-rest/internal/util"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
)

func Authenticate(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")
	if token == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	email, userId, err := util.VerifyToken(token)
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	ctx.Set("email", email)
	ctx.Set("userId", userId)
	ctx.Next()
}

func AuthorizeEventOwnerEdit(ctx *gin.Context) {
	id := ctx.Param("id")
	event, err := models.GetEventById(id)
	if err != nil {
		logger.Logger.Warning(err)
		ctx.AbortWithStatusJSON(http.StatusNotFound, models.NewErrMessage(fmt.Sprintf("id not found - %s", id)))
		return
	}
	userId, _ := strconv.ParseInt(ctx.GetString("userId"), 10, 64)
	if event.UserId != userId {
		logger.Logger.Warning(err)
		ctx.AbortWithStatusJSON(http.StatusForbidden, models.NewErrMessage(fmt.Sprintf("not owner of event - %s", id)))
		return
	}
	ctx.Next()
}

// LogrusHttp logging middleware for every HTTP request/response
func LogrusHttp(logger log.FieldLogger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		raw := ctx.Request.URL.RawQuery
		start := time.Now()
		// Process request
		ctx.Next()
		// Fill the params
		param := gin.LogFormatterParams{}
		param.TimeStamp = time.Now() // Stop timer
		param.Latency = param.TimeStamp.Sub(start)
		if param.Latency > time.Minute {
			param.Latency = param.Latency.Truncate(time.Second)
		}

		param.ClientIP = ctx.ClientIP()
		param.Method = ctx.Request.Method
		param.StatusCode = ctx.Writer.Status()
		param.BodySize = ctx.Writer.Size()
		if raw != "" {
			path = path + "?" + raw
		}
		param.Path = path

		fields := map[string]any{
			"client_id":   param.ClientIP,
			"method":      param.Method,
			"status_code": param.StatusCode,
			"body_size":   param.BodySize,
			"path":        param.Path,
			"latency":     param.Latency.String(),
		}

		entry := logger.WithFields(fields)

		if ctx.Writer.Status() >= 500 {
			entry.Error()
		} else {
			entry.Info()
		}
	}
}
