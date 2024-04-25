package snippetshandler

import (
	"github.com/labstack/echo/v4"
	"github.com/rsvix/go-htmx-app-template/internal/templates"
)

type getSnippetDeleteModalEditHandlerParams struct {
	pageTitle string
}

func GetSnippetDeleteModalEditHandler() *getSnippetDeleteModalEditHandlerParams {
	return &getSnippetDeleteModalEditHandlerParams{
		pageTitle: "Snippets",
	}
}

func (h getSnippetDeleteModalEditHandlerParams) Serve(c echo.Context) error {
	snippetId := c.Param("id")
	return templates.DeleteModal(snippetId).Render(c.Request().Context(), c.Response())
}
