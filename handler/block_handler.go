package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/ehardi19/rantaiblok/model"
	"github.com/labstack/echo"
)

// IsValid handles validity function
func (h *Handler) IsValid(c echo.Context) error {
	valid, err := h.Service.IsValid()

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"valid": valid,
	})
}

// SaveBlock handles saving block
func (h *Handler) SaveBlock(c echo.Context) error {
	var req model.CreateBlockRequest

	body := c.Request().Body
	err := json.NewDecoder(body).Decode(&req)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	err = h.Service.SaveBlock(req)

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, echo.Map{"created_at": time.Now()})
}

// GetAllBlock handles getting all block
func (h *Handler) GetAllBlock(c echo.Context) error {
	blocks, err := h.Service.GetAllBlock()

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, blocks)
}

// GetLastBlock handles getting last block
func (h *Handler) GetLastBlock(c echo.Context) error {
	block, err := h.Service.GetLastBlock()

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, block)
}

// GetBlockByID handles getting an block by id
func (h *Handler) GetBlockByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	block, err := h.Service.GetBlockByID(id)

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, block)
}

// PushDataToBlock handles publish block from data pool request
func (h *Handler) PushDataToBlock(c echo.Context) error {
	err := h.Service.PushDataToBlock()

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, echo.Map{"created_at": time.Now()})
}
