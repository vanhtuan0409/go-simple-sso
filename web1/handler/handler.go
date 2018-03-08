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

type homeViewModel struct {
	Name string
}

func (h *handler) Home(c echo.Context) error {
	name, _ := c.Cookie("name")
	data := homeViewModel{
		Name: name.Value,
	}
	return c.Render(http.StatusOK, "home.html", data)
}
