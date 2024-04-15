package handlers

import (
	"encoding/hex"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type GetProcessHandler struct {
}

func NewGetProcessHandler() *GetProcessHandler {
	return &GetProcessHandler{}
}

func (i GetProcessHandler) ServeHTTP(c echo.Context) error {

	token := c.Param("token")
	log.Printf("token: %s\n", token)

	if !strings.Contains(token, "O") {
		return c.Redirect(http.StatusSeeOther, "/login")
	}

	// Get session
	session, err := session.Get("authenticate-sessions", c)
	if err != nil {
		log.Printf("Error getting session: %v\n", err)
	}

	// Check if session is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		validator := strings.Split(token, "O")[0]
		extraInfo := strings.Split(token, "O")[1]
		log.Printf("validator: %s\n", validator)
		log.Printf("extraInfo: %s\n", extraInfo)

		decoded, err := hex.DecodeString(extraInfo)
		if err != nil {
			log.Println(err)
			return err
		}
		decodedStr := string(decoded[:])
		mode := strings.Split(decodedStr, "@")[0]
		id := strings.Split(decodedStr, "@")[1]
		log.Printf("mode: %s\n", mode)
		log.Printf("id: %s\n", id)

		if mode == "activate" {
			log.Println("Processing activation")

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
					if strings.Compare(strings.TrimSpace(result.Activationtoken), strings.TrimSpace(validator)) == 0 {
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

		} else if mode == "resetpwd" {
			log.Println("Processing password reset")
		}
		return nil
	}
	return c.Redirect(http.StatusSeeOther, "/")
}
