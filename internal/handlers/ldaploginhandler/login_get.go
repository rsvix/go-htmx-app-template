package ldaploginhandler

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/rsvix/go-htmx-app-template/internal/templates"
)

type getLdapLoginHandlerParams struct {
	appName   string
	pageTitle string
}

func GetLdapLoginHandler() *getLdapLoginHandlerParams {
	return &getLdapLoginHandlerParams{
		appName:   os.Getenv("APP_NAME"),
		pageTitle: "Login",
	}
}

func (h getLdapLoginHandlerParams) Serve(c echo.Context) error {
	csrfToken := ""
	// If you prefer to use TokenLookup: "form:_csrf" in your CSRF middleware
	// if value, ok := c.Get(middleware.DefaultCSRFConfig.ContextKey).(string); ok {
	// 	csrfToken = value
	// }

	return templates.LdapLoginPage(c, h.appName, h.pageTitle, csrfToken).Render(c.Request().Context(), c.Response())
}
