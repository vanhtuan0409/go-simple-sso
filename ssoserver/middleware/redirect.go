package middleware

import (
	"github.com/labstack/echo"
	"github.com/vanhtuan0409/go-simple-sso/ssoserver/config"
	"github.com/vanhtuan0409/go-simple-sso/ssoserver/handler"
)

func RedirectCallbackMiddleware(env *config.AppEnv) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cookie, err := c.Cookie("session_id")
			if err != nil || cookie.Value == "" {
				return next(c)
			}

			session, err := env.Datastore.GetSession(cookie.Value)
			if err != nil {
				return next(c)
			}

			return handler.RedirectCallback(c, env, session.Token)
		}
	}
}
