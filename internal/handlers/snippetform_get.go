package handlers

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/rsvix/go-htmx-app-template/internal/templates"
)

type getSnippetFormHandlerParams struct {
	pageTitle string
}

func GetSnippetFormHandler() *getSnippetFormHandlerParams {
	return &getSnippetFormHandlerParams{
		pageTitle: "Snippets",
	}
}

func (h getSnippetFormHandlerParams) Serve(c echo.Context) error {
	snippetId := c.Param("id")
	log.Println(snippetId)

	var languages = []string{"bash",
		"c",
		"cpp",
		"csharp",
		"css",
		"diff",
		"go",
		"graphql",
		"ini",
		"java",
		"javascript",
		"json",
		"kotlin",
		"less",
		"lua",
		"makefile",
		"markdown",
		"objectivec",
		"pearl",
		"php",
		"php-template",
		"python",
		"python-repl",
		"r",
		"ruby",
		"rust",
		"scss",
		"shell",
		"sql",
		"swift",
		"typescript",
		"vbnet",
		"wasm",
		"xml",
		"yaml",
	}

	return templates.SnippetFormModal(languages).Render(c.Request().Context(), c.Response())
}
