package snippetshandler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/rsvix/go-htmx-app-template/internal/structs"
	"gorm.io/gorm"
)

type postSnippetFormHandlerParams struct {
	pageTitle string
}

func PostSnippetFormHandler() *postSnippetFormHandlerParams {
	return &postSnippetFormHandlerParams{
		pageTitle: "Snippets",
	}
}

func (h postSnippetFormHandlerParams) Serve(c echo.Context) error {

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
	userIdInt, err := strconv.ParseUint(userIdString, 10, 32)
	if err != nil {
		fmt.Println(err)
		return c.HTML(http.StatusInternalServerError, "<h2>Error, please try again</h2>")
	}
	userIdUint := uint(userIdInt)

	var userName string
	if value, ok := session.Values["firstname"].(string); ok {
		userName = value
	} else {
		return c.HTML(http.StatusInternalServerError, "<h2>Error, please try again</h2>")
	}

	snippetName := c.Request().FormValue("snippetName")
	log.Println(snippetName)

	snippetLanguage := c.Request().FormValue("snippetLanguage")
	log.Println(snippetLanguage)

	snippetContent := c.Request().FormValue("snippetContent")
	log.Println(snippetContent)

	publicFlag := c.Request().FormValue("publicSnippet")
	log.Println(publicFlag)
	var publicSnippet uint = 0
	if publicFlag == "true" {
		publicSnippet = 1
	}

	currentUrl := c.Request().Header.Get("HX-Current-URL")
	log.Println(currentUrl)

	snippet := structs.Snippet{
		Owner:     userIdUint,
		Ownername: userName,
		Name:      snippetName,
		Language:  snippetLanguage,
		Code:      snippetContent,
		Ispublic:  publicSnippet,
	}

	db := c.Get("__db").(*gorm.DB)
	result := db.Create(&snippet)
	if err := result.Error; err != nil {
		log.Println(err.Error())
		return c.HTML(http.StatusInternalServerError, "<h2>Error, please try again</h2>")
	}

	return c.Redirect(http.StatusSeeOther, "/snippets")
}
