package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/ehardi19/rantaiblok/model"
	"github.com/labstack/echo"
)

// SaveAkta ...
func (h *Handler) SaveAkta(c echo.Context) error {
	var akta model.Akta

	body := c.Request().Body
	err := json.NewDecoder(body).Decode(&akta)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	err = h.Service.SaveAkta(akta)

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, echo.Map{"created_at": time.Now()})
}

// GetAllAkta ...
func (h *Handler) GetAllAkta(c echo.Context) error {
	aktas, err := h.Service.GetAllAkta()

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, aktas)
}

// GetAktaByID ...
func (h *Handler) GetAktaByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	akta, err := h.Service.GetAktaByID(id)

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, akta)
}

// DeleteAktaByID ...
func (h *Handler) DeleteAktaByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	err = h.Service.DeleteAktaByID(id)

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, nil)
}
