package main

import (
	"fmt"
	"html/template"
	"io"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/vanhtuan0409/go-simple-sso/web/config"
	"github.com/vanhtuan0409/go-simple-sso/web/handler"
	mdw "github.com/vanhtuan0409/go-simple-sso/web/middleware"
	"github.com/vanhtuan0409/go-simple-sso/web/service"
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
	cfg := config.Parse()

	// Dependencies
	s := service.NewTokenVerifyService(cfg.VERIFY_TOKEN_URL)
	t := newTpl("template/*.html")

	// App env
	appEnv := new(config.AppEnv)
	appEnv.Config = cfg
	appEnv.TokenVerifyService = s

	// Handler
	h := new(handler.Handler)
	h.AppEnv = appEnv

	// Routing
	e := echo.New()
	e.Renderer = t
	authMiddleware := mdw.AuthMiddleware(appEnv)

	e.Use(middleware.Logger())
	e.GET("/", h.Home, authMiddleware)
	e.GET("/callback", h.Callback)
	e.Start(fmt.Sprintf(":%d", cfg.HTTP_PORT))
}
