package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/rsvix/go-htmx-app-template/internal/structs"
	"github.com/rsvix/go-htmx-app-template/internal/templates"
)

type getSnippetsHandlerParams struct {
	pageTitle string
}

func GetSnippetsHandler() *getSnippetsHandlerParams {
	return &getSnippetsHandlerParams{
		pageTitle: "Snippets",
	}
}

func (h getSnippetsHandlerParams) Serve(c echo.Context) error {

	snippets := []structs.Snippet{
		{
			Id:       1,
			Name:     "test snippet",
			Category: "public",
			Language: "golang",
			Owner:    "me",
			Content:  "asdasd",
		},
		{
			Id:       2,
			Name:     "test snippet 2",
			Category: "public",
			Language: "rust",
			Owner:    "you",
			Content:  "asdascascasd",
		},
		{
			Id:       3,
			Name:     "test snippet 3",
			Category: "private",
			Language: "python",
			Owner:    "john",
			Content:  "1231231232",
		},
	}
	// log.Println(snippets)

	userName := c.Get("userName").(string)
	return templates.SnippetsPage(c, h.pageTitle, userName, snippets).Render(c.Request().Context(), c.Response())
}
