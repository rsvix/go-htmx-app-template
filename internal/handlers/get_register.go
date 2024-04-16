package handlers

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/rsvix/go-htmx-app-template/internal/templates"
)

type getRegisterHandlerParams struct {
	appName   string
	pageTitle string
}

func GetRegisterHandler() *getRegisterHandlerParams {
	return &getRegisterHandlerParams{
		appName:   os.Getenv("APP_NAME"),
		pageTitle: "Register",
	}
}

func (h getRegisterHandlerParams) Serve(c echo.Context) error {
	// session, err := session.Get("authenticate-sessions", c)
	// if err != nil {
	// 	log.Printf("Error getting session: %v\n", err)
	// }
	// if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
	// 	return templates.RegisterPage(h.appName, h.pageTitle).Render(c.Request().Context(), c.Response())
	// }
	// return c.Redirect(http.StatusSeeOther, "/")

	return templates.RegisterPage(h.appName, h.pageTitle).Render(c.Request().Context(), c.Response())
}
