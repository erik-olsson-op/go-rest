package main

import (
	"github.com/erik-olsson-op/go-rest/internal/database"
	"github.com/erik-olsson-op/go-rest/internal/server"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apiKey JWT
// @in header
// @name Authorization

func main() {
	database.Init()
	server.Init()
}
