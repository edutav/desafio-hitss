package routes

import (
	"net/http"

	"github.com/edutav/golang-api-clients/config/database"
	"github.com/edutav/golang-api-clients/docs"
	"github.com/edutav/golang-api-clients/internal/controllers"
	"github.com/edutav/golang-api-clients/internal/repositories"
	"github.com/edutav/golang-api-clients/internal/services"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func Init() *echo.Echo {
	e := echo.New()

	//TODO Swagger info
	docs.SwaggerInfo.Title = "Clients API"
	docs.SwaggerInfo.Description = "Essa documentação é para o desafio da Hitss"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	group := e.Group("/api/v1")

	group.GET("/", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "Olá Echo Framework!")
	})

	db := database.CreateConnection()

	mq := services.NewRabbitMQ()

	repo := repositories.NewClientsRepository(db)
	clientsController := controllers.NewClientsController(*repo, *mq)

	group.GET("/clients", clientsController.GetAll)
	group.GET("/clients/:id", clientsController.GetByID)
	group.POST("/clients", clientsController.Store)
	group.PUT("/clients/:id", clientsController.Update)
	group.DELETE("/clients/:id", clientsController.Delete)

	return e
}
