package handlers

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/rsvix/go-htmx-app-template/internal/templates"
)

type getSnippetHandlerParams struct {
	pageTitle string
}

func GetSnippetHandler() *getSnippetHandlerParams {
	return &getSnippetHandlerParams{
		pageTitle: "Snippets",
	}
}

func (h getSnippetHandlerParams) Serve(c echo.Context) error {
	snippetId := c.Param("id")
	log.Println(snippetId)
	return templates.SnippetModal(c, h.pageTitle, snippetId).Render(c.Request().Context(), c.Response())
}
