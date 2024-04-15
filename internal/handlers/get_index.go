package handlers

import (
	"log"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/rsvix/go-htmx-app-template/internal/templates"
)

type GetIndexHandler struct {
	pageTitle string
}

func NewGetIndexHandler() *GetIndexHandler {
	return &GetIndexHandler{
		pageTitle: "Index",
	}
}

func (h GetIndexHandler) ServeHTTP(c echo.Context) error {

	session, err := session.Get("authenticate-sessions", c)
	if err != nil {
		log.Printf("Get: %v\n", err)
	}

	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		return c.Redirect(http.StatusSeeOther, "/login")
		// return templates.LoginPage("Login", false, "").Render(c.Request().Context(), c.Response())
	}

	userName := "user"
	if value, ok := session.Values["firstname"].(string); ok {
		userName = value
	}
	return templates.IndexPage(h.pageTitle, userName).Render(c.Request().Context(), c.Response())
}
