package server

import (
	_ "github.com/erik-olsson-op/go-rest/docs"
	"github.com/erik-olsson-op/go-rest/internal/handlers"
	"github.com/erik-olsson-op/go-rest/internal/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterHealthCheck(server *gin.Engine) {
	server.GET("/api/ping", handlers.Ping)
}

func RegisterSwagger(server *gin.Engine) {
	server.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func RegisterV1Routes(server *gin.Engine) {

	v1Protected := server.Group("api/v1")
	{
		v1Protected.Use(middleware.Authenticate)
		// events
		v1Protected.POST("/events", handlers.CreateEventV1)
		v1Protected.DELETE("/events/:id", middleware.AuthorizeEventOwnerEdit, handlers.DeleteEventByIdV1)
		v1Protected.PUT("/events/:id", middleware.AuthorizeEventOwnerEdit, handlers.UpdateEventByIdV1)
	}

	v1 := server.Group("api/v1")
	{
		// events
		v1.GET("/events", handlers.GetAllEventsV1)
		v1.GET("/events/:id", handlers.GetEventByIdV1)
		// users
		v1.POST("/users", handlers.SignUpV1)
		v1.POST("/users/login", handlers.LoginV1)
	}
}
