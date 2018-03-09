package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func (h *Handler) Logout(c echo.Context) error {
	cookie, _ := c.Cookie("session_id")
	h.AppEnv.Datastore.DeleteSession(cookie.Value)
	cookie.Value = ""
	c.SetCookie(cookie)
	return c.Redirect(http.StatusFound, "/")
}
