package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/rsvix/go-htmx-app-template/internal/templates"
)

type GetLoginHandler struct {
	appName   string
	pageTitle string
}

func NewGetLoginHandler() *GetLoginHandler {
	return &GetLoginHandler{
		appName:   os.Getenv("APP_NAME"),
		pageTitle: "Login",
	}
}

func (h GetLoginHandler) ServeHTTP(c echo.Context) error {

	session, err := session.Get("authenticate-sessions", c)
	if err != nil {
		log.Printf("Error getting session: %v\n", err)
	}

	// Check if session is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		return templates.LoginPage(h.appName, h.pageTitle).Render(c.Request().Context(), c.Response())
	}
	return c.Redirect(http.StatusSeeOther, "/")
}
