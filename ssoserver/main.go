package main

import (
	"html/template"
	"io"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/vanhtuan0409/go-simple-sso/ssoserver/config"
	"github.com/vanhtuan0409/go-simple-sso/ssoserver/datastore"
	"github.com/vanhtuan0409/go-simple-sso/ssoserver/handler"
	mdw "github.com/vanhtuan0409/go-simple-sso/ssoserver/middleware"
	"github.com/vanhtuan0409/go-simple-sso/ssoserver/model"
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
	ds := datastore.NewDatastore()
	ds.SaveUser(model.NewUser("member@pav.com", "abc123", "Member 1"))

	t := newTpl("template/*.html")

	// App env
	appEnv := new(config.AppEnv)
	appEnv.Config = cfg
	appEnv.Datastore = ds

	// Create handler
	h := new(handler.Handler)
	h.AppEnv = appEnv

	// Routing
	e := echo.New()
	e.Use(middleware.Logger())
	e.Renderer = t

	redirectMdw := mdw.RedirectCallbackMiddleware(appEnv)

	e.GET("/", h.LoginView, redirectMdw)
	e.POST("/", h.LoginProcess, redirectMdw)
	e.GET("/logout", h.Logout)
	e.POST("/verify_token", h.VerifyToken)
	e.Start(":5000")
}
