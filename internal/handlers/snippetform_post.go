package handlers

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
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
	snippetName := c.Request().FormValue("snippetName")
	log.Println(snippetName)

	snippetLanguage := c.Request().FormValue("snippetLanguage")
	log.Println(snippetLanguage)

	snippetContent := c.Request().FormValue("snippetContent")
	log.Println(snippetContent)

	publicSnippet := c.Request().FormValue("publicSnippet")
	log.Println(publicSnippet)

	currentUrl := c.Request().Header.Get("HX-Current-URL")
	log.Println(currentUrl)

	return c.Redirect(http.StatusSeeOther, "/snippets")
}
