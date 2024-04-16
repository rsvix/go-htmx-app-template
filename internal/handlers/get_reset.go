package handlers

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/rsvix/go-htmx-app-template/internal/templates"
)

type getResetHandlerParams struct {
	appName   string
	pageTitle string
}

func GetResetHandler() *getResetHandlerParams {
	return &getResetHandlerParams{
		appName:   os.Getenv("APP_NAME"),
		pageTitle: "Register",
	}
}

func (h getResetHandlerParams) Serve(c echo.Context) error {
	// session, err := session.Get("authenticate-sessions", c)
	// if err != nil {
	// 	log.Printf("Error getting session: %v\n", err)
	// }
	// if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
	// 	return templates.ResetPage(h.appName, h.pageTitle).Render(c.Request().Context(), c.Response())
	// }
	// return c.Redirect(http.StatusSeeOther, "/")

	return templates.ResetPage(h.appName, h.pageTitle).Render(c.Request().Context(), c.Response())
}
