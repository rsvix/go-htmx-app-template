package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type deleteSnippetEditHandlerParams struct {
	pageTitle string
}

func DeleteSnippetEditHandler() *deleteSnippetEditHandlerParams {
	return &deleteSnippetEditHandlerParams{
		pageTitle: "Snippets",
	}
}

func (h deleteSnippetEditHandlerParams) Serve(c echo.Context) error {
	snippetId := c.Param("id")

	db := c.Get("__db").(*gorm.DB)
	var result struct {
		Name     string
		Language string
		Code     string
	}
	db.Raw("DELETE FROM snippets WHERE id = ?;", snippetId).Scan(&result)

	return c.Redirect(http.StatusSeeOther, "/snippets")
}
