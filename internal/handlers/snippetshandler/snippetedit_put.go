package snippetshandler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
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
	snippetId := c.Param("id")
	snippetContent := c.Request().FormValue("snippetContent")
	currentUrl := c.Request().Header.Get("HX-Current-URL")
	log.Println(currentUrl)

	db := c.Get("__db").(*gorm.DB)
	var result struct {
		Name     string
		Language string
		Code     string
	}
	db.Raw("UPDATE snippets SET code = ? WHERE id = ?;", snippetContent, snippetId).Scan(&result)

	return c.Redirect(http.StatusSeeOther, "/snippets")
}
