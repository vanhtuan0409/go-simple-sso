package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/vanhtuan0409/go-simple-sso/web/config"
)

func RedirectToLogin(c echo.Context, env *config.AppEnv) error {
	callbackURL := env.Config.SERVER_URL + "/callback"
	loginURL := env.Config.SSO_URL + "?callback=" + callbackURL
	return c.Redirect(http.StatusFound, loginURL)
}
