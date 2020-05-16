package handler

import (
	"errors"
	"net/http"

	"github.com/ehardi19/rantaiblok/service"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

// Handler ..
type Handler struct {
	Service service.Service
}

// InitHandler ..
func InitHandler(service service.Service) Handler {
	h := Handler{service}

	return h
}

// HelloWorld ..
func (h *Handler) HelloWorld(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, echo.Map{
		"hello": "world",
	})
}

// ResponseError ..
type ResponseError struct {
	Message string `json:"message"`
}

var (
	// ErrInternalServerError will throw if any the Internal Server Error happen
	ErrInternalServerError = errors.New("Internal Server Error")
	// ErrNotFound will throw if the requested item is not exists
	ErrNotFound = errors.New("Your requested Item is not found")
	// ErrConflict will throw if the current action already exists
	ErrConflict = errors.New("Your Item already exist")
	// ErrBadParamInput will throw if the given request-body or params is not valid
	ErrBadParamInput = errors.New("Given Param is not valid")
)

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	logrus.Error(err)
	switch err {
	case ErrInternalServerError:
		return http.StatusInternalServerError
	case ErrNotFound:
		return http.StatusNotFound
	case ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
