package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/vanhtuan0409/go-simple-sso/web1/handler"
)

const (
	SSO_ADDRESS    = "http://login.com:5000"
	SERVER_ADDRESS = "http://web1.com:8081"
)

type tpl struct {
	templates *template.Template
}

func newTpl(pattern string) *tpl {
	return &tpl{
		templates: template.Must(template.ParseGlob(pattern)),
	}
}

func (t *tpl) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func authMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if cookie, err := c.Cookie("name"); err != nil || cookie.Value == "" {
			callbackURL := SERVER_ADDRESS + "/callback"
			loginURL := SSO_ADDRESS + "?callback=" + callbackURL
			return c.Redirect(http.StatusFound, loginURL)
		}

		return next(c)
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	// Set golang template
	t := newTpl("template/*.html")
	e.Renderer = t

	// Routing
	e.GET("/", handler.Home, authMiddleware)
	e.GET("/callback", handler.Callback)
	e.GET("/logout", handler.Logout)
	e.Start(":8081")
}
