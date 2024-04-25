package activationhandler

import (
	"log"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/rsvix/go-htmx-app-template/internal/templates"
)

type activateHandlerParams struct {
	pageTitle string
}

func GetActivateHandler() *activateHandlerParams {
	return &activateHandlerParams{
		pageTitle: "Activate",
	}
}

func (h activateHandlerParams) Serve(c echo.Context) error {
	session, err := session.Get("authenticate-sessions", c)
	if err != nil {
		log.Printf("Error getting session: %v\n", err)
		return c.Redirect(http.StatusSeeOther, "/error")
	}

	if value, ok := session.Values["user_enabled"].(string); ok {
		if value == "1" {
			return templates.ActivatePage(c, h.pageTitle, true, "Account activated").Render(c.Request().Context(), c.Response())
		} else if value == "2" {
			return templates.ActivatePage(c, h.pageTitle, false, "Invalid token").Render(c.Request().Context(), c.Response())
		} else if value == "3" {
			return templates.ActivatePage(c, h.pageTitle, false, "Token expired").Render(c.Request().Context(), c.Response())
		}
	}
	return c.Redirect(http.StatusSeeOther, "/login")
}
