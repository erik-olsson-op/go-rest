package server

import (
	"github.com/erik-olsson-op/go-rest/internal/logger"
	"github.com/erik-olsson-op/go-rest/internal/middleware"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		logger.Logger.Warning("failed to read .env file")
	}
	_, addrOk := os.LookupEnv("SERVER_ADDR")
	if !addrOk {
		logger.Logger.Fatal("SERVER_ADDR not set")
	}
}

// Init the Gin web server
func Init() {
	mode := gin.DebugMode
	gin.SetMode(mode)
	server := gin.New()
	err := server.SetTrustedProxies(nil)
	if err != nil {
		logger.Logger.Fatalln("failed to read .env file")
	}
	server.Use(gin.Recovery())
	server.Use(middleware.LogrusHttp(logger.Logger))
	RegisterHealthCheck(server)
	RegisterSwagger(server)
	RegisterV1Routes(server)
	err = server.Run(os.Getenv("SERVER_ADDR"))
	if err != nil {
		logger.Logger.Fatalln("failed to read .env file")
	}
}
