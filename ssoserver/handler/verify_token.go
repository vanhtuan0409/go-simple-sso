package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

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

func renderVerifyTokenError(c echo.Context, err error) error {
	return c.JSON(http.StatusBadRequest, map[string]interface{}{
		"success": false,
		"message": err.Error(),
	})
}
