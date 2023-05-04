package main

import (
	"github.com/edutav/golang-api-clients/config/database"
	"github.com/edutav/golang-api-clients/internal/routes"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func main() {
	database.Init()

	echo := routes.Init()

	echo.Logger.Fatal(echo.Start(":8080"))
}
