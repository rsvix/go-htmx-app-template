package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/rsvix/go-htmx-app-template/internal/templates"
)

type getSnippetDeleteModalEditHandlerParams struct {
	pageTitle string
}

func GetSnippetDeleteModalEditHandler() *getSnippetDeleteModalEditHandlerParams {
	return &getSnippetDeleteModalEditHandlerParams{
		pageTitle: "Snippets",
	}
}

func (h getSnippetDeleteModalEditHandlerParams) Serve(c echo.Context) error {
	snippetId := c.Param("id")

	// session, err := session.Get("authenticate-sessions", c)
	// if err != nil {
	// 	log.Printf("Error getting session: %v\n", err)
	// 	return c.HTML(http.StatusInternalServerError, "<h2>Error, please try again</h2>")
	// }

	// var userIdString string
	// if value, ok := session.Values["id"].(string); ok {
	// 	userIdString = value
	// } else {
	// 	return c.HTML(http.StatusInternalServerError, "<h2>Error, please try again</h2>")
	// }
	// log.Println(userIdString)

	// db := c.Get("__db").(*gorm.DB)

	// var owner string
	// db.Raw("SELECT owner FROM snippets WHERE id = ?;", snippetId).Scan(&owner)
	// log.Println(owner)

	return templates.DeleteModal(snippetId).Render(c.Request().Context(), c.Response())

	// if strings.Compare(strings.TrimSpace(owner), strings.TrimSpace(userIdString)) == 0 {
	// 	db.Exec("DELETE FROM snippets WHERE id = ?;", snippetId)
	// }

	// return c.Redirect(http.StatusSeeOther, "/snippets")
}
