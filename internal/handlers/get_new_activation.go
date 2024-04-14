package handlers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type GetNewActivationHandler struct {
}

func NewGetNewActivationHandler() *GetNewActivationHandler {
	return &GetNewActivationHandler{}
}

func (i GetNewActivationHandler) ServeHTTP(c echo.Context) error {

	// Get session
	session, err := session.Get("authenticate-sessions", c)
	if err != nil {
		log.Printf("Error getting session: %v\n", err)
	}

	id := session.Values["id"].(string)
	log.Printf("id: %v\n", id)

	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		db := c.Get("__db").(*gorm.DB)
		var result struct {
			Email                     string
			Activationtokenexpiration time.Time
			Enabled                   int
		}
		db.Raw("SELECT email, activationtokenexpiration, enabled FROM users WHERE id = ?", id).Scan(&result)
		log.Printf("result: %v\n", result)

		if result.Enabled == 0 {
			diff := result.Activationtokenexpiration.Sub(time.Now().UTC())
			secs := diff.Seconds()
			log.Printf("diff: %v\nsecs: %v\n", diff, secs)

			if secs > 0.0 {
				return c.HTML(http.StatusInternalServerError, "<h2>You already have a valid<br/>activation link in your email</h2>")
			}
			// send new email
			return c.HTML(http.StatusInternalServerError, fmt.Sprintf("<h2>Activation link sent to<br/>%s</h2>", result.Email))
		}
		return c.HTML(http.StatusInternalServerError, "<h2>Couldn't process your request</h2>")
	}
	return c.Redirect(http.StatusSeeOther, "/")
}
