package snippetshandler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rsvix/go-htmx-app-template/internal/utils"
	"gorm.io/gorm"
)

type putSnippetEditHandlerParams struct {
	pageTitle string
}

func PutSnippetEditHandler() *putSnippetEditHandlerParams {
	return &putSnippetEditHandlerParams{
		pageTitle: "Snippets",
	}
}

func (h putSnippetEditHandlerParams) Serve(c echo.Context) error {
	sessionInfo, err := utils.GetSessionInfo(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	snippetId := c.Param("id")
	snippetContent := c.Request().FormValue("snippetContent")
	currentUrl := c.Request().Header.Get("HX-Current-URL")
	log.Println(currentUrl)

	db := c.Get("__db").(*gorm.DB)
	var owner int
	_ = db.Raw("SELECT owner FROM snippets WHERE id = ?;", snippetId).Scan(&owner)
	if owner == sessionInfo.Id {
		var result struct {
			Name     string
			Language string
			Code     string
		}
		db.Raw("UPDATE snippets SET code = ? WHERE id = ?;", snippetContent, snippetId).Scan(&result)
	}
	return c.Redirect(http.StatusSeeOther, "/snippets")
}
