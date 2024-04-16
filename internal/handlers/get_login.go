package handlers

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rsvix/go-htmx-app-template/internal/templates"
)

type getLoginHandlerParams struct {
	appName   string
	pageTitle string
}

func GetLoginHandler() *getLoginHandlerParams {
	return &getLoginHandlerParams{
		appName:   os.Getenv("APP_NAME"),
		pageTitle: "Login",
	}
}

func (h getLoginHandlerParams) Serve(c echo.Context) error {
	// session, err := session.Get("authenticate-sessions", c)
	// if err != nil {
	// 	log.Printf("Error getting session: %v\n", err)
	// }
	// if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
	// 	return templates.LoginPage(h.appName, h.pageTitle).Render(c.Request().Context(), c.Response())
	// }
	// return c.Redirect(http.StatusSeeOther, "/")

	csrfToken := "none"
	if value, ok := c.Get(middleware.DefaultCSRFConfig.ContextKey).(string); ok {
		csrfToken = value
	}
	return templates.LoginPage(c, h.appName, h.pageTitle, csrfToken).Render(c.Request().Context(), c.Response())
}
