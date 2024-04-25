package registerhandler

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	// To use TokenLookup: "form:_csrf"
	csrfToken := "none"
	if value, ok := c.Get(middleware.DefaultCSRFConfig.ContextKey).(string); ok {
		csrfToken = value
	}
	return templates.RegisterPage(c, h.appName, h.pageTitle, csrfToken).Render(c.Request().Context(), c.Response())
}
