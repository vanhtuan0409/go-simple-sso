package main

import (
	"html/template"
	"io"

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

func main() {
	e := echo.New()

	// Set golang template
	t := newTpl("template/*.html")
	e.Renderer = t

	// Create handler
	h := handler.NewHandler()

	// Routing
	e.GET("/", h.Home)
	e.Start(":8081")
}
