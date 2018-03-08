package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

type handler struct {
}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) Home(c echo.Context) error {
	return c.String(http.StatusOK, "home")
}
