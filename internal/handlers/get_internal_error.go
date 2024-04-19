package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/rsvix/go-htmx-app-template/internal/templates"
)

type getInternalErrorHandlerParams struct {
	pageTitle string
}

func GetInternalErrorHandler() *getInternalErrorHandlerParams {
	return &getInternalErrorHandlerParams{
		pageTitle: "NotFound",
	}
}

func (h getInternalErrorHandlerParams) Serve(c echo.Context) error {
	return templates.ErrorPage(c, "Error", "We are working to fix the problem").Render(c.Request().Context(), c.Response())
}
