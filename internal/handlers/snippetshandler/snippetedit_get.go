package snippetshandler

import (
	"github.com/labstack/echo/v4"
	"github.com/rsvix/go-htmx-app-template/internal/templates"
	"gorm.io/gorm"
)

type getSnippetEditHandlerParams struct {
	pageTitle string
}

func GetSnippetEditHandler() *getSnippetEditHandlerParams {
	return &getSnippetEditHandlerParams{
		pageTitle: "Snippets",
	}
}

func (h getSnippetEditHandlerParams) Serve(c echo.Context) error {
	snippetId := c.Param("id")

	db := c.Get("__db").(*gorm.DB)
	var result struct {
		Name     string
		Language string
		Code     string
	}
	db.Raw("SELECT name, language, code FROM snippets WHERE id = ?;", snippetId).Scan(&result)

	return templates.SnippetEditModal(snippetId, result.Name, result.Language, result.Code).Render(c.Request().Context(), c.Response())
}
