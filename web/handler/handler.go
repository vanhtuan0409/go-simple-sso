package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/vanhtuan0409/go-simple-sso/web/config"
	"github.com/vanhtuan0409/go-simple-sso/web/model"
)

type handler struct {
	appEnv *config.AppEnv
}

func NewHandler(env *config.AppEnv) *handler {
	return &handler{
		appEnv: env,
	}
}

type homeViewModel struct {
	User       *model.User
	Title      string
	LogoutLink string
}

func (h *handler) Home(c echo.Context) error {
	data := c.Get("user")
	user, ok := data.(*model.User)
	if !ok {
		return c.String(http.StatusInternalServerError, "Some error happen")
	}

	viewData := homeViewModel{
		User:       user,
		Title:      h.appEnv.Config.APP_TITLE,
		LogoutLink: h.appEnv.Config.SSS_URL + "/logout",
	}
	return c.Render(http.StatusOK, "home.html", viewData)
}

func (h *handler) Callback(c echo.Context) error {
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
