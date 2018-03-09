package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/vanhtuan0409/go-simple-sso/ssoserver/config"
)

func RedirectCallback(c echo.Context, cfg *config.Config, token string) error {
	callback := c.Request().URL.Query().Get("callback")
	if callback == "" {
		callback = cfg.DefaultCallback
	}
	tokenAddedCallback := callback + "?token=" + token
	return c.Redirect(http.StatusFound, tokenAddedCallback)
}
