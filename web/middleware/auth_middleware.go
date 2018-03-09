package middleware

import (
	"github.com/labstack/echo"
	"github.com/vanhtuan0409/go-simple-sso/web/config"
	"github.com/vanhtuan0409/go-simple-sso/web/handler"
	"github.com/vanhtuan0409/go-simple-sso/web/service"
)

func AuthMiddleware(cfg *config.Config, s service.TokenVerifyService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cookie, err := c.Cookie("token")
			if err != nil || cookie.Value == "" {
				return handler.RedirectToLogin(c, cfg)
			}

			user, err := s.Verify(cookie.Value)
			if err != nil {
				cookie.Value = ""
				c.SetCookie(cookie)
				return handler.RedirectToLogin(c, cfg)
			}

			c.Set("user", user)
			return next(c)
		}
	}
}
