package accounthandler

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/rsvix/go-htmx-app-template/internal/templates"
)

type putEditAccountHandlerParams struct {
	pageTitle string
	dbKey     string
}

func PutEditAccountHandler() *putEditAccountHandlerParams {
	return &putEditAccountHandlerParams{
		pageTitle: "Account",
		dbKey:     os.Getenv("DB_CONTEXT_KEY"),
	}
}

func (h putEditAccountHandlerParams) Serve(c echo.Context) error {
	email := c.Request().FormValue("email")
	firstname := c.Request().FormValue("firstname")
	lastname := c.Request().FormValue("lastname")

	return templates.EditAccountConfirmModal(email, firstname, lastname).Render(c.Request().Context(), c.Response())
}
