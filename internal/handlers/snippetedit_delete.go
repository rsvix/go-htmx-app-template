package handlers

import (
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo-contrib/session"
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

	session, err := session.Get("authenticate-sessions", c)
	if err != nil {
		log.Printf("Error getting session: %v\n", err)
		return c.HTML(http.StatusInternalServerError, "<h2>Error, please try again</h2>")
	}

	var userIdString string
	if value, ok := session.Values["id"].(string); ok {
		userIdString = value
	} else {
		return c.HTML(http.StatusInternalServerError, "<h2>Error, please try again</h2>")
	}
	log.Println(userIdString)

	db := c.Get("__db").(*gorm.DB)

	var owner string
	db.Raw("SELECT owner FROM snippets WHERE id = ?;", snippetId).Scan(&owner)
	log.Println(owner)

	if strings.Compare(strings.TrimSpace(owner), strings.TrimSpace(userIdString)) == 0 {
		db.Exec("DELETE FROM snippets WHERE id = ?;", snippetId)
	}

	return c.Redirect(http.StatusSeeOther, "/snippets")
}
