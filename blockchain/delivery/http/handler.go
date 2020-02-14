package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"

	"github.com/ehardi19/rantaiblok/blockchain"
	"github.com/ehardi19/rantaiblok/models"
)

// ResponseError ..
type ResponseError struct {
	Message string `json:"message"`
}

// Handler ..
type Handler struct {
	Usecase blockchain.Usecase
}

// NewHandler ..
func NewHandler(e *echo.Echo, usecase blockchain.Usecase) {
	handler := &Handler{
		Usecase: usecase,
	}
	e.GET("/", handler.HelloWorld)
	e.GET("/blockchain", handler.Fetch)
	e.GET("/blockchain/:id", handler.GetByID)
	e.POST("/blockchain", handler.Store)
	e.GET("/blockchain/valid", handler.Validate)
}

// HelloWorld ..
func (h *Handler) HelloWorld(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"hello": "world",
	})
}

// Fetch ..
func (h *Handler) Fetch(c echo.Context) error {
	blockchain, err := h.Usecase.Fetch()

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, blockchain)
}

// GetByID ..
func (h *Handler) GetByID(c echo.Context) error {
	ids, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, models.ErrNotFound.Error())
	}

	id := int64(ids)

	block, err := h.Usecase.GetByID(id)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, block)
}

// Store ..
func (h *Handler) Store(c echo.Context) error {
	var req models.BlockRequest

	body := c.Request().Body
	err := json.NewDecoder(body).Decode(&req)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	block, err := h.Usecase.Store(req)

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, block)
}

// Validate ..
func (h *Handler) Validate(c echo.Context) error {
	status := h.Usecase.Validate()

	if !status {
		return c.JSON(http.StatusOK, echo.Map{
			"status": "invalid",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status": "valid",
	})
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	logrus.Error(err)
	switch err {
	case models.ErrInternalServerError:
		return http.StatusInternalServerError
	case models.ErrNotFound:
		return http.StatusNotFound
	case models.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
