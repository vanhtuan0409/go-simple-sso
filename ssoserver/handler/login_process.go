package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/vanhtuan0409/go-simple-sso/ssoserver/model"
)

type loginRequest struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

func (h *Handler) LoginProcess(c echo.Context) error {
	request := new(loginRequest)
	if err := c.Bind(request); err != nil {
		return renderLoginProcessError(c, "Request error")
	}

	user, err := h.AppEnv.Datastore.GetUserByEmail(request.Email)
	if err != nil {
		return renderLoginProcessError(c, "Email or Password is incorrect")
	}

	if !user.CheckPassword(request.Password) {
		return renderLoginProcessError(c, "Email or Password is incorrect")
	}

	session := model.NewSession(user.ID)
	if err := h.AppEnv.Datastore.SaveSession(session); err != nil {
		return renderLoginProcessError(c, "Internal error happened")
	}

	cookie := &http.Cookie{
		Name:  "session_id",
		Value: session.ID,
	}
	c.SetCookie(cookie)
	return RedirectCallback(c, h.AppEnv, session.Token)
}

func renderLoginProcessError(c echo.Context, err string) error {
	data := loginViewModel{}
	data.Error = err
	return c.Render(http.StatusOK, "login.html", data)
}
