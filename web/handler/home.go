package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/vanhtuan0409/go-simple-sso/web/model"
)

type homeViewModel struct {
	User       *model.User
	Title      string
	LogoutLink string
}

func (h *Handler) Home(c echo.Context) error {
	data := c.Get("user")
	user, ok := data.(*model.User)
	if !ok {
		return c.String(http.StatusInternalServerError, "Some error happen")
	}

	viewData := homeViewModel{
		User:       user,
		Title:      h.AppEnv.Config.APP_TITLE,
		LogoutLink: h.AppEnv.Config.SSS_URL + "/logout",
	}
	return c.Render(http.StatusOK, "home.html", viewData)
}
