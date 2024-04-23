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
			Id:        1,
			Name:      "test snippet",
			Ispublic:  1,
			Language:  "golang",
			Owner:     1,
			Ownername: "asdfasd",
			Code:      "asdasd",
		},
		{
			Id:        2,
			Name:      "test snippet 2",
			Ispublic:  1,
			Language:  "rust",
			Owner:     1,
			Ownername: "asdfasd",
			Code:      "asdascascasd",
		},
		{
			Id:        3,
			Name:      "test snippet 3",
			Ispublic:  0,
			Language:  "python",
			Owner:     2,
			Ownername: "asdfasd",
			Code:      "1231231232",
		},
	}
	// log.Println(snippets)

	userName := c.Get("userName").(string)
	return templates.SnippetsPage(c, h.pageTitle, userName, snippets).Render(c.Request().Context(), c.Response())
}
