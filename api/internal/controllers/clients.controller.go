package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/edutav/golang-api-clients/internal/models"
	"github.com/edutav/golang-api-clients/internal/repositories"
	"github.com/edutav/golang-api-clients/internal/services"
	"github.com/edutav/golang-api-clients/pkg/utils"
	"github.com/labstack/echo/v4"
)

type ClientsControllerStruct struct {
	clientsRepository *repositories.ClientsRepository
	Mq                *services.RabbitMQ
}

type ClientsController interface {
	GetAll(ctx echo.Context)
	Store(ctx echo.Context)
	GetByID(ctx echo.Context)
	Update(ctx echo.Context)
	Delete(ctx echo.Context)
}

func NewClientsController(clientsRepository repositories.ClientsRepository, mq services.RabbitMQ) *ClientsControllerStruct {
	return &ClientsControllerStruct{
		clientsRepository: &clientsRepository,
		Mq:                &mq,
	}
}

// GetAll Obtém uma lista de clientes
// @Summary Obtém uma lista de clientes
// @Description Obtém uma lista de clientes existente
// @Tags clients
// @Accept json
// @Produce json
// @Success 200 {object} models.Clients
// @Failure 404 {object} utils.HttpError
// @Failure 500 {object} utils.HttpError
// @Router /clients [get]
func (c *ClientsControllerStruct) GetAll(ctx echo.Context) error {
	result, err := c.clientsRepository.GetAll()
	if err != nil {
		utils.NewError(ctx, http.StatusInternalServerError, err)
		return err
	}
	return ctx.JSON(http.StatusOK, result)
}

// Store cria um novo cliente
// @Summary Cria um novo cliente
// @Description Cria um novo cliente com base nos dados fornecidos
// @Tags clients
// @Accept json
// @Produce json
// @Param cliente body models.Clients true "Client ADD"
// @Success 201 {object} models.Clients
// @Failure 500 {object} utils.HttpError
// @Router /clients [post]
func (c *ClientsControllerStruct) Store(ctx echo.Context) error {
	client := new(models.Clients)
	if err := ctx.Bind(client); err != nil {
		utils.NewError(ctx, http.StatusBadRequest, err)
		return err
	}

	message, err := json.Marshal(client)
	if err != nil {
		utils.NewError(ctx, http.StatusInternalServerError, err)
		return err
	}

	// mq := services.NewRabbitMQ()
	if err := c.Mq.SendMessage(message); err != nil {
		utils.NewError(ctx, http.StatusInternalServerError, err)
		return err
	}

	return ctx.JSON(http.StatusOK, nil)
}

// GetByID obtém um cliente pelo ID
// @Summary Obtém um cliente pelo ID
// @Description Obtém um cliente existente com base no ID fornecido
// @Tags clients
// @Accept json
// @Produce json
// @Param id path int true "ID do cliente"
// @Success 200 {object} models.Clients
// @Failure 400 {object} utils.HttpError
// @Router /clients/{id} [get]
func (c *ClientsControllerStruct) GetByID(ctx echo.Context) error {
	clientID := ctx.Param("id")
	ID, err := strconv.Atoi(clientID)
	if err != nil {
		utils.NewError(ctx, http.StatusBadRequest, err)
		return err
	}

	client, err := c.clientsRepository.GetByID(ID)
	if err != nil {
		utils.NewError(ctx, http.StatusInternalServerError, err)
		return err
	}

	return ctx.JSON(http.StatusOK, client)
}

// Update atualiza um cliente existente
// @Summary Atualiza um cliente existente
// @Description Atualiza um cliente existente com base no ID fornecido
// @Tags clients
// @Accept json
// @Produce json
// @Param id path int true "ID do cliente"
// @Param cliente body models.Clients true "Dados do cliente"
// @Success 204 ""
// @Failure 400 {object} utils.HttpError
// @Failure 500 {object} utils.HttpError
// @Router /produtos/{id} [put]
func (c *ClientsControllerStruct) Update(ctx echo.Context) error {
	clientID := ctx.Param("id")
	ID, err := strconv.Atoi(clientID)
	if err != nil {
		utils.NewError(ctx, http.StatusBadRequest, err)
		return err
	}

	updateClient := new(models.Clients)
	if err := ctx.Bind(updateClient); err != nil {
		utils.NewError(ctx, http.StatusBadRequest, err)
		return err
	}

	if err := c.clientsRepository.Update(ID, updateClient); err != nil {
		utils.NewError(ctx, http.StatusInternalServerError, err)
		return err
	}

	return ctx.JSON(http.StatusNoContent, nil)
}

// Delete exclui um cliente pelo ID
// @Summary Exclui um cliente pelo ID
// @Description Exclui um cliente existente com base no ID fornecido
// @Tags clients
// @Accept json
// @Produce json
// @Param id path int true "ID do cliente"
// @Success 204 ""
// @Failure 400 {object} utils.HttpError
// @Failure 500 {object} utils.HttpError
// @Router /produtos/{id} [delete]
func (c *ClientsControllerStruct) Delete(ctx echo.Context) error {
	clientID := ctx.Param("id")
	ID, err := strconv.Atoi(clientID)
	if err != nil {
		utils.NewError(ctx, http.StatusBadRequest, err)
		return err
	}

	if err := c.clientsRepository.Delete(ID); err != nil {
		utils.NewError(ctx, http.StatusInternalServerError, err)
		return err
	}

	return ctx.JSON(http.StatusNoContent, nil)
}
