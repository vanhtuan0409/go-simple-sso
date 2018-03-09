package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/vanhtuan0409/go-simple-sso/web/config"
)

func RedirectToLogin(c echo.Context, cfg *config.Config) error {
	callbackURL := cfg.SERVER_URL + "/callback"
	loginURL := cfg.SSS_URL + "?callback=" + callbackURL
	return c.Redirect(http.StatusFound, loginURL)
}
