package snippetshandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rsvix/go-htmx-app-template/internal/templates"
	"github.com/rsvix/go-htmx-app-template/internal/utils"
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
	sessionInfo, err := utils.GetSessionInfo(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	snippetId := c.Param("id")
	db := c.Get("__db").(*gorm.DB)
	var owner int
	_ = db.Raw("SELECT owner FROM snippets WHERE id = ?;", snippetId).Scan(&owner)
	if owner == sessionInfo.Id {
		var result struct {
			Name     string
			Language string
			Code     string
		}
		_ = db.Raw("SELECT name, language, code FROM snippets WHERE id = ?;", snippetId).Scan(&result)
		return templates.SnippetEditModal(snippetId, result.Name, result.Language, result.Code).Render(c.Request().Context(), c.Response())
	}
	return c.Redirect(http.StatusSeeOther, "/snippets")
}
