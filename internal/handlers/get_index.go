package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/rsvix/go-htmx-app-template/internal/templates"
)

type getIndexHandlerParams struct {
	pageTitle string
}

func GetIndexHandler() *getIndexHandlerParams {
	return &getIndexHandlerParams{
		pageTitle: "Index",
	}
}

func (h getIndexHandlerParams) Serve(c echo.Context) error {
	// session, err := session.Get("authenticate-sessions", c)
	// if err != nil {
	// 	log.Printf("Error getting session: %v\n", err)
	// }
	// if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
	// 	return c.Redirect(http.StatusSeeOther, "/login")
	// 	// return templates.LoginPage("Login", false, "").Render(c.Request().Context(), c.Response())
	// }
	// userName := "User"
	// if value, ok := session.Values["firstname"].(string); ok {
	// 	userName = value
	// }
	// return templates.IndexPage(h.pageTitle, userName).Render(c.Request().Context(), c.Response())

	userName := c.Get("userName").(string)
	return templates.IndexPage(c, h.pageTitle, userName).Render(c.Request().Context(), c.Response())
}
