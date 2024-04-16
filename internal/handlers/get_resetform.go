package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/rsvix/go-htmx-app-template/internal/templates"
)

type getResetformHandlerParams struct {
	appName   string
	pageTitle string
}

func GetResetformHandler() *getResetformHandlerParams {
	return &getResetformHandlerParams{
		appName:   os.Getenv("APP_NAME"),
		pageTitle: "Reset",
	}
}

func (h getResetformHandlerParams) Serve(c echo.Context) error {

	session, err := session.Get("authenticate-sessions", c)
	if err != nil {
		log.Printf("Session error: %v\n", err)
	}

	if session.Values["pwreset"] != nil {
		if auth, ok := session.Values["pwreset"].(bool); !ok || !auth {
			en_err := session.Values["en_error"].(string)
			return templates.ResetFormPage(c, h.appName, h.pageTitle, false, "", en_err).Render(c.Request().Context(), c.Response())
		}
		id := session.Values["user_id"].(string)
		return templates.ResetFormPage(c, h.appName, h.pageTitle, true, id, "Reset your password").Render(c.Request().Context(), c.Response())
	}
	return c.Redirect(http.StatusSeeOther, "/login")
}
