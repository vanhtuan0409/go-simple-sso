package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/vanhtuan0409/go-simple-sso/web2/model"
)

const (
	SSO_ADDRESS = "http://login.com:5000"
)

type homeViewModel struct {
	User       *model.User
	LogoutLink string
}

func Home(c echo.Context) error {
	data := c.Get("user")
	user, ok := data.(*model.User)
	if !ok {
		return c.String(http.StatusInternalServerError, "Some error happen")
	}

	viewData := homeViewModel{
		User:       user,
		LogoutLink: SSO_ADDRESS + "/logout",
	}
	return c.Render(http.StatusOK, "home.html", viewData)
}

func Callback(c echo.Context) error {
	token := c.Request().URL.Query().Get("token")
	if token == "" {
		return c.String(http.StatusBadRequest, "Token is required")
	}

	cookie := &http.Cookie{
		Name:  "token",
		Value: token,
	}
	c.SetCookie(cookie)
	return c.Redirect(http.StatusFound, "/")
}
