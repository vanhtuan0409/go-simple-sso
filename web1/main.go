package main

import (
	"html/template"
	"io"
	"net/http"
	"net/url"

	"github.com/labstack/echo"
	"github.com/vanhtuan0409/go-simple-sso/web1/handler"
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
		if _, err := c.Cookie("name"); err != nil {
			u, _ := url.Parse("http://localhost:5000")
			query := url.Values{}
			requestURL := c.Request().URL
			requestURL.Host = "http://localhost:8081"
			query.Set("callback", requestURL.String())
			u.RawQuery = query.Encode()

			return c.Redirect(http.StatusTemporaryRedirect, u.String())
		}

		return nil
	}
}

func main() {
	e := echo.New()

	// Set golang template
	t := newTpl("template/*.html")
	e.Renderer = t

	// Routing
	e.GET("/", handler.Home, authMiddleware)
	e.Start(":8081")
}
