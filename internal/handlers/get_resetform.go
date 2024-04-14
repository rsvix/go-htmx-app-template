package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/rsvix/go-htmx-app-template/templates"
)

type GetResetformHandler struct {
}

func NewGetResetformHandler() *GetResetformHandler {
	return &GetResetformHandler{}
}

func (i GetResetformHandler) ServeHTTP(c echo.Context) error {
	appName := os.Getenv("APP_NAME")

	session, err := session.Get("authenticate-sessions", c)
	if err != nil {
		log.Printf("Session error: %v\n", err)
	}

	if session.Values["pwreset"] != nil {
		if auth, ok := session.Values["pwreset"].(bool); !ok || !auth {
			en_err := session.Values["en_error"].(string)
			return templates.ResetFormPage(appName, "Reset", false, "", en_err).Render(c.Request().Context(), c.Response())
		}
		id := session.Values["user_id"].(string)
		return templates.ResetFormPage(appName, "Reset", true, id, "Reset your password").Render(c.Request().Context(), c.Response())
	}
	return c.Redirect(http.StatusSeeOther, "/login")
}
