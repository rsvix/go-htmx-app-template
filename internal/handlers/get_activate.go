package handlers

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
	}

	if session.Values["enabled"] != nil {
		if auth, ok := session.Values["enabled"].(bool); !ok || !auth {
			// id := session.Values["id"].(string)
			en_err := session.Values["en_error"].(string)
			// return templates.ActivatePage("Activate", false, id, en_err).Render(c.Request().Context(), c.Response())
			return templates.ActivatePage(c, h.pageTitle, false, en_err).Render(c.Request().Context(), c.Response())
		}
		// return templates.ActivatePage("Activate", true, "", "Account activated").Render(c.Request().Context(), c.Response())
		return templates.ActivatePage(c, h.pageTitle, true, "Account activated").Render(c.Request().Context(), c.Response())
	}
	return c.Redirect(http.StatusSeeOther, "/login")
}
