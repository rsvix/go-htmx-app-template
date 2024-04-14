package handlers

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type GetActivationHandler struct {
}

func NewGetActivationHandler() *GetActivationHandler {
	return &GetActivationHandler{}
}

func (i GetActivationHandler) ServeHTTP(c echo.Context) error {

	// Query params
	maskedId := c.Param("id")
	activationToken := c.Param("activationtoken")

	log.Printf("maskedId: %s\n", maskedId)
	s, err := strconv.Atoi(maskedId)
	if err != nil {
		log.Printf("Error converting value: %s\n", err)
		return nil
	}
	id := strconv.Itoa((s - 575426791))
	log.Printf("id: %s\n", id)

	// Get session
	session, err := session.Get("authenticate-sessions", c)
	if err != nil {
		log.Printf("Error getting session: %v\n", err)
	}

	session.Values["enabled"] = false
	session.Values["id"] = id

	db := c.Get("__db").(*gorm.DB)
	var result struct {
		Activationtoken           string
		Activationtokenexpiration time.Time
		Enabled                   int
	}
	db.Raw("SELECT activationtoken, activationtokenexpiration, enabled FROM users WHERE id = ?", id).Scan(&result)

	if result.Enabled == 0 {
		// diff := time.Until(result.Activationtokenexpiration)
		diff := result.Activationtokenexpiration.Sub(time.Now().UTC())
		secs := diff.Seconds()
		// log.Printf("diff: %v\nsecs: %v\n", diff, secs)

		if secs > 0.0 {
			if strings.Compare(strings.TrimSpace(result.Activationtoken), strings.TrimSpace(activationToken)) == 0 {
				// res := db.Table("users").Where("id = ?", id).Update("enabled", 1)
				timeNow := time.Now().UTC()
				res := db.Table("users").Where("id = ?", id).Updates(map[string]interface{}{"enabled": 1, "activationtokenexpiration": timeNow})
				if res.Error != nil {
					log.Println("Error enabling user")
					return nil
				}
				// log.Println("User enabled")
				session.Values["enabled"] = true
				if err := session.Save(c.Request(), c.Response()); err != nil {
					log.Printf("Error saving session: %s", err)
					return nil
				}
				return c.Redirect(http.StatusSeeOther, "/activate")
			}
			session.Values["en_error"] = "Invalid token"
			err = session.Save(c.Request(), c.Response())
			if err != nil {
				log.Printf("Error saving session: %v\n", err)
			}
			return c.Redirect(http.StatusSeeOther, "/activate")
		}
		session.Values["en_error"] = "Token expired"
		err = session.Save(c.Request(), c.Response())
		if err != nil {
			log.Printf("Error saving session: %v\n", err)
		}
		return c.Redirect(http.StatusSeeOther, "/activate")
	}
	session.Values["en_error"] = "User already enabled"
	err = session.Save(c.Request(), c.Response())
	if err != nil {
		log.Printf("Error saving session: %v\n", err)
	}
	return c.Redirect(http.StatusSeeOther, "/activate")
}
