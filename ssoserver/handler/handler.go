package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/vanhtuan0409/go-simple-sso/ssoserver/datastore"
)

type handler struct {
	ds datastore.Datastore
}

func NewHandler(ds datastore.Datastore) *handler {
	return &handler{
		ds: ds,
	}
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
