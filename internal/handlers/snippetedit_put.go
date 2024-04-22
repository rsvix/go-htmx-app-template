package handlers

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
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
	snippetContent := c.Request().FormValue("snippetContent")
	log.Println(snippetContent)

	currentUrl := c.Request().Header.Get("HX-Current-URL")
	log.Println(currentUrl)

	return c.Redirect(http.StatusSeeOther, "/snippets")
}
