package activationhandler

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type getNewActivationHandler struct {
}

func GetNewActivationHandler() *getNewActivationHandler {
	return &getNewActivationHandler{}
}

func (h getNewActivationHandler) Serve(c echo.Context) error {

	session, err := session.Get("authenticate-sessions", c)
	if err != nil {
		log.Printf("Error getting session: %v\n", err)
	}

	// id := session.Values["user_id"].(string)
	var id string
	if value, ok := session.Values["user_id"].(string); !ok {
		return c.Redirect(http.StatusSeeOther, "/login")
	} else {
		id = value
	}
	log.Printf("id: %v\n", id)

	if auth, _ := session.Values["authenticated"].(bool); !auth {
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
