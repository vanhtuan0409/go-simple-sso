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

type loginViewModel struct {
}

func (h *handler) LoginView(c echo.Context) error {
	data := loginViewModel{}
	return c.Render(http.StatusOK, "login.html", data)
}

func (h *handler) LoginProcess(c echo.Context) error {
	data := loginViewModel{}
	return c.Render(http.StatusOK, "login.html", data)
}
