package snippetshandler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/rsvix/go-htmx-app-template/internal/structs"
	"github.com/rsvix/go-htmx-app-template/internal/utils"
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

	sessionInfo, err := utils.GetSessionInfo(c)
	if err != nil {
		log.Printf("Error getting session info: %v\n", err)
		return c.Redirect(http.StatusSeeOther, "/error")
	}

	userIdInt, err := strconv.ParseUint(sessionInfo.Id, 10, 32)
	if err != nil {
		fmt.Println(err)
		return c.Redirect(http.StatusSeeOther, "/error")
	}

	snippetName := c.Request().FormValue("snippetName")
	snippetLanguage := c.Request().FormValue("snippetLanguage")
	snippetContent := c.Request().FormValue("snippetContent")
	publicFlag := c.Request().FormValue("publicSnippet")
	var publicSnippet uint = 0
	if publicFlag == "true" {
		publicSnippet = 1
	}

	currentUrl := c.Request().Header.Get("HX-Current-URL")
	log.Println(currentUrl)

	snippet := structs.Snippet{
		Owner:     uint(userIdInt),
		Ownername: sessionInfo.Username,
		Name:      snippetName,
		Language:  snippetLanguage,
		Code:      snippetContent,
		Ispublic:  publicSnippet,
	}

	db := c.Get("__db").(*gorm.DB)
	result := db.Create(&snippet)
	if err := result.Error; err != nil {
		log.Println(err.Error())
		return c.Redirect(http.StatusSeeOther, "/error")
	}
	return c.Redirect(http.StatusSeeOther, "/snippets")
}
