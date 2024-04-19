package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/rsvix/go-htmx-app-template/internal/templates"
)

type getNotfoundHandlerParams struct {
	pageTitle string
}

func GetNotfoundHandler() *getNotfoundHandlerParams {
	return &getNotfoundHandlerParams{
		pageTitle: "NotFound",
	}
}

func (h getNotfoundHandlerParams) Serve(c echo.Context) error {
	return templates.NotfoundPage(c, "Not Found", "Sorry, we can't find that page").Render(c.Request().Context(), c.Response())
}
