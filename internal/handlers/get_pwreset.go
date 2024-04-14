package handlers

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type GetPwResetHandler struct {
}

func NewGetPwResetHandler() *GetPwResetHandler {
	return &GetPwResetHandler{}
}

func (i GetPwResetHandler) ServeHTTP(c echo.Context) error {

	id := c.Param("id")
	resetToken := c.Param("resettoken")

	session, err := session.Get("authenticate-sessions", c)
	if err != nil {
		log.Printf("Get: %v\n", err)
	}

	session.Values["pwreset"] = false
	session.Values["pr_error"] = ""
	session.Values["user_id"] = ""

	db := c.Get("__db").(*gorm.DB)

	var result struct {
		Email                         string
		Passwordchangetoken           string
		Passwordchangetokenexpiration time.Time
		Enabled                       int
	}
	db.Raw("SELECT email, passwordchangetoken, passwordchangetokenexpiration, enabled FROM users WHERE id = ?", id).Scan(&result)

	// diff := time.Until(result.Passwordchangetokenexpiration)
	diff := result.Passwordchangetokenexpiration.Sub(time.Now().UTC())
	secs := diff.Seconds()
	// log.Printf("%v\n%v", diff, secs)

	if secs < 0.0 {
		return c.JSON(http.StatusBadRequest, "Token expired")
		// session.Values["pr_error"] = "Token expired"
		// err = session.Save(c.Request(), c.Response())
		// if err != nil {
		// 	log.Printf("Error saving session: %v\n", err)
		// }
		// return c.Redirect(http.StatusSeeOther, "/resetform")
	}

	if result.Enabled == 0 {
		return c.JSON(http.StatusBadRequest, "User must be activated first")
		// session.Values["pr_error"] = "User must be activated first"
		// err = session.Save(c.Request(), c.Response())
		// if err != nil {
		// 	log.Printf("Error saving session: %v\n", err)
		// }
		// return c.Redirect(http.StatusSeeOther, "/resetform")
	}

	log.Printf("resetToken: %v\n", resetToken)
	log.Printf("Passwordchangetoken: %v\n", result.Passwordchangetoken)

	if strings.Compare(strings.TrimSpace(result.Passwordchangetoken), strings.TrimSpace(resetToken)) == 0 {
		session.Values["pwreset"] = true
		session.Values["user_id"] = id
		err = session.Save(c.Request(), c.Response())
		if err != nil {
			log.Printf("Error saving session: %v\n", err)
		}
		return c.Redirect(http.StatusSeeOther, "/resetform")
	}

	return c.JSON(http.StatusBadRequest, "Invalid token")
	// session.Values["pr_error"] = "Invalid token"
	// err = session.Save(c.Request(), c.Response())
	// if err != nil {
	// 	log.Printf("Error saving session: %v\n", err)
	// }
	// return c.Redirect(http.StatusSeeOther, "/resetform")
}
