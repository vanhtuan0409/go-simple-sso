package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/vanhtuan0409/go-simple-sso/ssoserver/config"
)

func RedirectCallback(c echo.Context, env *config.AppEnv, token string) error {
	callback := c.Request().URL.Query().Get("callback")
	if callback == "" {
		callback = env.Config.DefaultCallback
	}
	tokenAddedCallback := callback + "?token=" + token
	return c.Redirect(http.StatusFound, tokenAddedCallback)
}
