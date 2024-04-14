package handlers

import (
	"log"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/rsvix/go-htmx-app-template/internal/templates"
)

type IndexHandler struct {
}

func (i IndexHandler) ServeHTTP(c echo.Context) error {

	// Get session
	session, err := session.Get("authenticate-sessions", c)
	if err != nil {
		log.Printf("Get: %v\n", err)
	}

	// Check if session is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		return c.Redirect(http.StatusSeeOther, "/login")
		// return templates.LoginPage("Login", false, "").Render(c.Request().Context(), c.Response())
	}

	// Get user firstName from session
	userName := "user"
	if value, ok := session.Values["firstname"].(string); ok {
		userName = value
	}
	return templates.IndexPage("index", userName).Render(c.Request().Context(), c.Response())
}
