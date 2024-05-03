package snippetshandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rsvix/go-htmx-app-template/internal/templates"
	"github.com/rsvix/go-htmx-app-template/internal/utils"
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
	sessionInfo, err := utils.GetSessionInfo(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	snippetId := c.Param("id")

	db := c.Get("__db").(*gorm.DB)
	var r1 struct {
		Owner    int
		Ispublic int
	}
	_ = db.Raw("SELECT owner, ispublic FROM snippets WHERE id = ?;", snippetId).Scan(&r1)
	if r1.Owner == sessionInfo.Id || r1.Ispublic == 1 {
		var result struct {
			Name     string
			Language string
			Code     string
		}
		db.Raw("SELECT name, language, code FROM snippets WHERE id = ?;", snippetId).Scan(&result)
		return templates.SnippetViewModal(result.Name, result.Language, result.Code).Render(c.Request().Context(), c.Response())
	}
	return c.Redirect(http.StatusSeeOther, "/snippets")
}
