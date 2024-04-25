package snippetshandler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rsvix/go-htmx-app-template/internal/utils"
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
	sessionInfo, err := utils.GetSessionInfo(c)
	if err != nil {
		log.Println(err)
		c.Response().Header().Set("HX-Redirect", "/error")
		return c.NoContent(http.StatusSeeOther)
	}

	snippetId := c.Param("id")
	db := c.Get("__db").(*gorm.DB)
	var owner int
	db.Raw("SELECT owner FROM snippets WHERE id = ?;", snippetId).Scan(&owner)
	log.Println(owner)

	if owner == sessionInfo.Id {
		db.Exec("DELETE FROM snippets WHERE id = ?;", snippetId)
	}

	return c.Redirect(http.StatusSeeOther, "/snippets")
}
