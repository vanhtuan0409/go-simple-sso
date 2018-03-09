package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/vanhtuan0409/go-simple-sso/web/handler"
	"github.com/vanhtuan0409/go-simple-sso/web/model"
)

const (
	SSO_ADDRESS    = "http://login.com:5000"
	SERVER_ADDRESS = "http://web.com:8081"
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

func redirectToLogin(c echo.Context) error {
	callbackURL := SERVER_ADDRESS + "/callback"
	loginURL := SSO_ADDRESS + "?callback=" + callbackURL
	return c.Redirect(http.StatusFound, loginURL)
}

type verifyResponse struct {
	Success bool        `json:"success"`
	User    *model.User `json:"user"`
	Message string      `json:"message"`
}

func verifyToken(token string) (*model.User, error) {
	verifyURL := SSO_ADDRESS + "/verify_token"

	data, err := json.Marshal(map[string]string{
		"token": token,
	})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", verifyURL, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	jsonObj := new(verifyResponse)
	if err := json.NewDecoder(resp.Body).Decode(jsonObj); err != nil {
		return nil, err
	}
	if !jsonObj.Success {
		return nil, errors.New(jsonObj.Message)
	}

	return jsonObj.User, nil
}

func authMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("token")
		if err != nil || cookie.Value == "" {
			return redirectToLogin(c)
		}

		user, err := verifyToken(cookie.Value)
		if err != nil {
			cookie.Value = ""
			c.SetCookie(cookie)
			return redirectToLogin(c)
		}

		c.Set("user", user)
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
	e.Start(":8081")
}
