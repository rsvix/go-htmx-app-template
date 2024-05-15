package scheduledhandler

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/rsvix/go-htmx-app-template/internal/templates"
)

type getNewJobHandlerParams struct {
	pageTitle string
}

func GetNewJobHandler() *getNewJobHandlerParams {
	return &getNewJobHandlerParams{
		pageTitle: "Snippets",
	}
}

func (h getNewJobHandlerParams) Serve(c echo.Context) error {
	snippetId := c.Param("id")
	log.Println(snippetId)

	var languages = []string{
		"bash",
		"c",
		"cpp",
		"csharp",
	}

	return templates.AddSchedJobModal(languages).Render(c.Request().Context(), c.Response())
}
