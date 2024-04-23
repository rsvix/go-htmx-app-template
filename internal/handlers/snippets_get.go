package handlers

import (
	"log"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/rsvix/go-htmx-app-template/internal/structs"
	"github.com/rsvix/go-htmx-app-template/internal/templates"
	"gorm.io/gorm"
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

	// snippets := []structs.Snippet{
	// 	{
	// 		Id:        1,
	// 		Name:      "test snippet",
	// 		Ispublic:  1,
	// 		Language:  "golang",
	// 		Owner:     1,
	// 		Ownername: "asdfasd",
	// 		Code:      "asdasd",
	// 	},
	// 	{
	// 		Id:        2,
	// 		Name:      "test snippet 2",
	// 		Ispublic:  1,
	// 		Language:  "rust",
	// 		Owner:     1,
	// 		Ownername: "asdfasd",
	// 		Code:      "asdascascasd",
	// 	},
	// 	{
	// 		Id:        3,
	// 		Name:      "test snippet 3",
	// 		Ispublic:  0,
	// 		Language:  "python",
	// 		Owner:     2,
	// 		Ownername: "asdfasd",
	// 		Code:      "1231231232",
	// 	},
	// }
	// log.Println(snippets)

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

	var userName string
	if value, ok := session.Values["firstname"].(string); ok {
		userName = value
	} else {
		return c.HTML(http.StatusInternalServerError, "<h2>Error, please try again</h2>")
	}

	db := c.Get("__db").(*gorm.DB)
	var result []structs.Snippet
	db.Raw("SELECT * FROM snippets WHERE owner = ? OR ispublic = ?;", userIdString, "1").Scan(&result)

	return templates.SnippetsPage(c, h.pageTitle, userName, result).Render(c.Request().Context(), c.Response())
}
