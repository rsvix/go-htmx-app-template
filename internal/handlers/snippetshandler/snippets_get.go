package snippetshandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rsvix/go-htmx-app-template/internal/structs"
	"github.com/rsvix/go-htmx-app-template/internal/templates"
	"github.com/rsvix/go-htmx-app-template/internal/utils"
	"gorm.io/gorm"
)

type getSnippetsHandlerParams struct {
	pageTitle string
}

func GetSnippetsHandler() *getSnippetsHandlerParams {
	return &getSnippetsHandlerParams{
		pageTitle: "Snippets",
	}
}

func (h getSnippetsHandlerParams) Serve(c echo.Context) error {
	sessionInfo, err := utils.GetSessionInfo(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	db := c.Get("__db").(*gorm.DB)
	var result []structs.Snippet
	_ = db.Raw("SELECT * FROM snippets WHERE owner = ? OR ispublic = ?;", sessionInfo.Id, "1").Scan(&result)

	return templates.SnippetsPage(c, h.pageTitle, sessionInfo.Username, result).Render(c.Request().Context(), c.Response())
}
