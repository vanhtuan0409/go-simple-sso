package middleware

import (
	"github.com/labstack/echo"
	"github.com/vanhtuan0409/go-simple-sso/web/config"
	"github.com/vanhtuan0409/go-simple-sso/web/handler"
)

func AuthMiddleware(env *config.AppEnv) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Redirect if there is no cookie
			cookie, err := c.Cookie("token")
			if err != nil || cookie.Value == "" {
				return handler.RedirectToLogin(c, env)
			}

			// Redirect if verification failed
			user, err := env.TokenVerifyService.Verify(cookie.Value)
			if err != nil {
				cookie.Value = ""
				c.SetCookie(cookie)
				return handler.RedirectToLogin(c, env)
			}

			c.Set("user", user)
			return next(c)
		}
	}
}
