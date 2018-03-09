package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func (h *Handler) Callback(c echo.Context) error {
	token := c.Request().URL.Query().Get("token")
	if token == "" {
		return c.String(http.StatusBadRequest, "Token is required")
	}

	cookie := &http.Cookie{
		Name:  "token",
		Value: token,
	}
	c.SetCookie(cookie)
	return c.Redirect(http.StatusFound, "/")
}
