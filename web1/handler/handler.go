package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

const (
	SSO_ADDRESS = "http://login.com:5000"
)

type homeViewModel struct {
	Name string
}

func Home(c echo.Context) error {
	name, _ := c.Cookie("name")
	data := homeViewModel{
		Name: name.Value,
	}
	return c.Render(http.StatusOK, "home.html", data)
}

func Callback(c echo.Context) error {
	name := c.Request().URL.Query().Get("name")
	if name == "" {
		return c.String(http.StatusBadRequest, "Name is required")
	}

	cookie := &http.Cookie{
		Name:  "name",
		Value: name,
	}
	c.SetCookie(cookie)
	return c.Redirect(http.StatusFound, "/")
}

func Logout(c echo.Context) error {
	cookie := &http.Cookie{
		Name:  "name",
		Value: "",
	}
	c.SetCookie(cookie)
	redirectURL := SSO_ADDRESS + "/logout"
	return c.Redirect(http.StatusFound, redirectURL)
}
