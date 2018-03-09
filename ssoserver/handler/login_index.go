package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

type loginViewModel struct {
	Error string
}

func (h *Handler) LoginView(c echo.Context) error {
	data := loginViewModel{}
	return c.Render(http.StatusOK, "login.html", data)
}
