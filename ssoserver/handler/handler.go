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
	Error string
}

func (h *handler) LoginView(c echo.Context) error {
	data := loginViewModel{}
	return c.Render(http.StatusOK, "login.html", data)
}

type loginRequest struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

func (h *handler) LoginProcess(c echo.Context) error {
	request := new(loginRequest)
	if err := c.Bind(request); err != nil {
		return renderLoginProcessError(c, "Request error")
	}

	user, err := h.ds.GetUserByEmail(request.Email)
	if err != nil {
		return renderLoginProcessError(c, "Email or Password is incorrect")
	}

	if !user.CheckPassword(request.Password) {
		return renderLoginProcessError(c, "Email or Password is incorrect")
	}

	callback := c.Request().URL.Query().Get("callback")
	return c.String(http.StatusOK, callback)
}

func renderLoginProcessError(c echo.Context, err string) error {
	data := loginViewModel{}
	data.Error = err
	return c.Render(http.StatusOK, "login.html", data)
}
