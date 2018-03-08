package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/vanhtuan0409/go-simple-sso/ssoserver/datastore"
	"github.com/vanhtuan0409/go-simple-sso/ssoserver/model"
)

type handler struct {
	ds datastore.Datastore
}

func NewHandler(ds datastore.Datastore) *handler {
	return &handler{
		ds: ds,
	}
}

type loginViewModel struct {
	Error string
}

func (h *handler) LoginView(c echo.Context) error {
	data := loginViewModel{}
	return c.Render(http.StatusOK, "login.html", data)
}

type loginRequest struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

func (h *handler) LoginProcess(c echo.Context) error {
	request := new(loginRequest)
	if err := c.Bind(request); err != nil {
		return renderLoginProcessError(c, "Request error")
	}

	user, err := h.ds.GetUserByEmail(request.Email)
	if err != nil {
		return renderLoginProcessError(c, "Email or Password is incorrect")
	}

	if !user.CheckPassword(request.Password) {
		return renderLoginProcessError(c, "Email or Password is incorrect")
	}

	session := model.NewSession(user.ID)
	if err := h.ds.SaveSession(session); err != nil {
		return renderLoginProcessError(c, "Internal error happened")
	}

	callback := c.Request().URL.Query().Get("callback")
	if callback == "" {
		callback = "http://web1.com:8081"
	}
	tokenAddedCallback := callback + "?token=" + session.Token

	cookie := &http.Cookie{
		Name:  "session_id",
		Value: session.ID,
	}
	c.SetCookie(cookie)
	return c.Redirect(http.StatusFound, tokenAddedCallback)
}

func (h *handler) Logout(c echo.Context) error {
	cookie, _ := c.Cookie("session_id")
	h.ds.DeleteSession(cookie.Value)
	cookie.Value = ""
	c.SetCookie(cookie)
	return c.Redirect(http.StatusFound, "/")
}

type verifyRequest struct {
	Token string `json:"token"`
}

func (h *handler) VerifyToken(c echo.Context) error {
	request := new(verifyRequest)
	if err := c.Bind(request); err != nil {
		c.Logger().Debugf("Error when parsing request: %v", err)
		return renderVerifyTokenError(c, err)
	}

	session, err := h.ds.GetSessionByToken(request.Token)
	if err != nil {
		c.Logger().Debugf("Error when getting session: %v", err)
		return renderVerifyTokenError(c, err)
	}

	user, err := h.ds.GetUser(session.UserID)
	if err != nil {
		c.Logger().Debugf("Error when getting user: %v", err)
		return renderVerifyTokenError(c, err)
	}

	fmt.Println("===========")
	fmt.Println("return user: ", user)
	fmt.Println("===========")

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"user":    user,
	})
}

func renderLoginProcessError(c echo.Context, err string) error {
	data := loginViewModel{}
	data.Error = err
	return c.Render(http.StatusOK, "login.html", data)
}

func renderVerifyTokenError(c echo.Context, err error) error {
	return c.JSON(http.StatusBadRequest, map[string]interface{}{
		"success": false,
		"message": err.Error(),
	})
}
