package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/rsvix/go-htmx-app-template/internal/templates"
	"gorm.io/gorm"
)

type getSnippetViewHandlerParams struct {
	pageTitle string
}

func GetSnippetViewHandler() *getSnippetViewHandlerParams {
	return &getSnippetViewHandlerParams{
		pageTitle: "Snippets",
	}
}

func (h getSnippetViewHandlerParams) Serve(c echo.Context) error {
	snippetId := c.Param("id")

	db := c.Get("__db").(*gorm.DB)
	var result struct {
		Name     string
		Language string
		Code     string
	}
	db.Raw("SELECT name, language, code FROM snippets WHERE id = ?;", snippetId).Scan(&result)

	return templates.SnippetViewModal(result.Name, result.Language, result.Code).Render(c.Request().Context(), c.Response())
}
