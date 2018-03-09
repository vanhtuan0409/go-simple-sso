package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/vanhtuan0409/go-simple-sso/ssoserver/config"
	"github.com/vanhtuan0409/go-simple-sso/ssoserver/model"
)

type Handler struct {
	AppEnv *config.AppEnv
}

type loginViewModel struct {
	Error string
}

func (h *Handler) LoginView(c echo.Context) error {
	data := loginViewModel{}
	return c.Render(http.StatusOK, "login.html", data)
}

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
	return RedirectCallback(c, h.AppEnv.Config, session.Token)
}

func (h *Handler) Logout(c echo.Context) error {
	cookie, _ := c.Cookie("session_id")
	h.AppEnv.Datastore.DeleteSession(cookie.Value)
	cookie.Value = ""
	c.SetCookie(cookie)
	return c.Redirect(http.StatusFound, "/")
}

type verifyRequest struct {
	Token string `json:"token"`
}

func (h *Handler) VerifyToken(c echo.Context) error {
	request := new(verifyRequest)
	if err := c.Bind(request); err != nil {
		c.Logger().Debugf("Error when parsing request: %v", err)
		return renderVerifyTokenError(c, err)
	}

	session, err := h.AppEnv.Datastore.GetSessionByToken(request.Token)
	if err != nil {
		c.Logger().Debugf("Error when getting session: %v", err)
		return renderVerifyTokenError(c, err)
	}

	user, err := h.AppEnv.Datastore.GetUser(session.UserID)
	if err != nil {
		c.Logger().Debugf("Error when getting user: %v", err)
		return renderVerifyTokenError(c, err)
	}

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
