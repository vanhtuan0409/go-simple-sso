package middleware

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/vanhtuan0409/go-simple-sso/ssoserver/datastore"
)

func RedirectMiddleware(ds datastore.Datastore) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cookie, err := c.Cookie("session_id")
			if err != nil || cookie.Value == "" {
				return next(c)
			}

			session, err := ds.GetSession(cookie.Value)
			if err != nil {
				return next(c)
			}

			callback := c.Request().URL.Query().Get("callback")
			if callback == "" {
				callback = "http://web1.com:8081"
			}
			tokenAddedCallback := callback + "?token=" + session.Token
			return c.Redirect(http.StatusFound, tokenAddedCallback)
		}
	}
}
