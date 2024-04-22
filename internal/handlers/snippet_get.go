package handlers

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/rsvix/go-htmx-app-template/internal/templates"
)

type getSnippetHandlerParams struct {
	pageTitle string
}

func GetSnippetHandler() *getSnippetHandlerParams {
	return &getSnippetHandlerParams{
		pageTitle: "Snippets",
	}
}

func (h getSnippetHandlerParams) Serve(c echo.Context) error {
	snippetId := c.Param("id")
	log.Println(snippetId)

	snippetName := "Test snippet"
	snippetLang := "go"
	snippetContent := `
package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
)

func main() {
	component := hello("John")
	
	http.Handle("/", templ.Handler(component))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
`

	return templates.SnippetModal(c, snippetName, snippetLang, snippetContent).Render(c.Request().Context(), c.Response())
}
