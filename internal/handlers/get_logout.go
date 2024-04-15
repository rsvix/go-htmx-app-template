package handlers

import (
	"log"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type logoutHandlerParams struct {
}

func GetLogoutHandler() *logoutHandlerParams {
	return &logoutHandlerParams{}
}

func (h logoutHandlerParams) Serve(c echo.Context) error {

	// Get session
	session, err := session.Get("authenticate-sessions", c)
	if err != nil {
		log.Printf("Error getting session: %v\n", err)
		return c.Redirect(http.StatusSeeOther, "/login")
	}

	// Check if session is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		return c.Redirect(http.StatusSeeOther, "/login")
	}

	// Set no cache response headers
	c.Response().Header().Set("Cache-Control", "no-cache, private, max-age=0")
	c.Response().Header().Set("Pragma", "no-cache")
	c.Response().Header().Set("X-Accel-Expires", "0")

	// Set session values
	session.Values["authenticated"] = false
	session.Values["email"] = nil
	session.Values["id"] = nil
	session.Values["firstname"] = nil
	session.Options.MaxAge = -1

	// Save updated session
	if err := session.Save(c.Request(), c.Response()); err != nil {
		log.Printf("Error saving session: %s", err)
	}

	return c.Redirect(http.StatusSeeOther, "/login")
}
