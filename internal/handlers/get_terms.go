package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/rsvix/go-htmx-app-template/internal/templates"
)

type getTermsHandlerParams struct {
	pageTitle string
}

func GetTermsHandlerParams() *getTermsHandlerParams {
	return &getTermsHandlerParams{
		pageTitle: "NotFound",
	}
}

func (h getTermsHandlerParams) Serve(c echo.Context) error {
	return templates.ErrorPage(c, "Error", "We are working to fix the problem").Render(c.Request().Context(), c.Response())
}
