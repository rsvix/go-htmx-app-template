package termshandler

import (
	"github.com/labstack/echo/v4"
	"github.com/rsvix/go-htmx-app-template/internal/templates"
)

type getTermsHandlerParams struct {
	pageTitle string
}

func GetTermsHandlerParams() *getTermsHandlerParams {
	return &getTermsHandlerParams{
		pageTitle: "Terms",
	}
}

func (h getTermsHandlerParams) Serve(c echo.Context) error {
	return templates.TermsPage(c, h.pageTitle).Render(c.Request().Context(), c.Response())
}
