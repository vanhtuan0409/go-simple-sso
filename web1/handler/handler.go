package handler

import (
	"net/http"

	"github.com/labstack/echo"
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
