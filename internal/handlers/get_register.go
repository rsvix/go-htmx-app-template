package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/rsvix/go-htmx-app-template/internal/templates"
)

type GetRegisterHandler struct {
}

func NewGetRegisterHandler() *GetLoginHandler {
	return &GetLoginHandler{}
}

func (i GetRegisterHandler) ServeHTTP(c echo.Context) error {

	// Get session
	session, err := session.Get("authenticate-sessions", c)
	if err != nil {
		log.Printf("Get: %v\n", err)
	}

	// Check if session is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		appName := os.Getenv("APP_NAME")
		return templates.RegisterPage(appName, "register").Render(c.Request().Context(), c.Response())
	}
	return c.Redirect(http.StatusSeeOther, "/")
}
