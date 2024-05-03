package loginhandler

import (
	"os"

	"github.com/labstack/echo/v4"
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
	csrfToken := ""
	// // To use TokenLookup: "form:_csrf"
	// if value, ok := c.Get(middleware.DefaultCSRFConfig.ContextKey).(string); ok {
	// 	csrfToken = value
	// }

	return templates.LoginPage(c, h.appName, h.pageTitle, csrfToken).Render(c.Request().Context(), c.Response())
}
