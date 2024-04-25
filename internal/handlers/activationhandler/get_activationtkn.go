package activationhandler

import (
	"encoding/hex"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type getActivationTokenHandlerParams struct {
	queryParam string
	dbKey      string
}

func GetActivationTokenHandler() *getActivationTokenHandlerParams {
	return &getActivationTokenHandlerParams{
		queryParam: "token",
		dbKey:      os.Getenv("DB_CONTEXT_KEY"),
	}
}

func (h getActivationTokenHandlerParams) Serve(c echo.Context) error {
	token := c.Param(h.queryParam)
	if !strings.Contains(token, "O") {
		return c.Redirect(http.StatusSeeOther, "/login")
	}

	extraInfo := strings.Split(token, "O")[1]
	decoded, err := hex.DecodeString(extraInfo)
	if err != nil {
		log.Println(err)
		return c.Redirect(http.StatusSeeOther, "/error")
	}
	decodedStr := string(decoded[:])
	mode := strings.Split(decodedStr, "@")[0]
	id := strings.Split(decodedStr, "@")[1]

	session, err := session.Get("authenticate-sessions", c)
	if err != nil {
		log.Printf("Error getting session: %v\n", err)
		return c.Redirect(http.StatusSeeOther, "/error")
	}

	if mode == "activate" {
		db := c.Get(h.dbKey).(*gorm.DB)
		var result struct {
			Activationtoken           string
			Activationtokenexpiration time.Time
			Enabled                   int
		}
		_ = db.Raw("SELECT activationtoken, activationtokenexpiration, enabled FROM users WHERE id = ?", id).Scan(&result)

		if result.Enabled == 0 {
			diff := result.Activationtokenexpiration.Sub(time.Now().UTC())
			secs := diff.Seconds()
			// log.Printf("diff: %v\nsecs: %v\n", diff, secs)

			if secs > 0.0 {
				if strings.Compare(strings.TrimSpace(result.Activationtoken), strings.TrimSpace(token)) == 0 {
					timeNow := time.Now().UTC()
					res := db.Table("users").Where("id = ?", id).Updates(map[string]interface{}{"enabled": 1, "activationtokenexpiration": timeNow})
					if res.Error != nil {
						log.Println("Error enabling user")
						return c.Redirect(http.StatusSeeOther, "/error")
					}
					session.Values["user_enabled"] = "1"
				} else {
					session.Values["user_enabled"] = "2"
					c.Set("enabling_error", "Invalid token")
				}
			} else {
				session.Values["user_enabled"] = "3"
				c.Set("enabling_error", "Token expired")
			}
		} else {
			return c.Redirect(http.StatusSeeOther, "/login")
		}
		if err := session.Save(c.Request(), c.Response()); err != nil {
			log.Printf("Error saving session: %s", err)
			return c.Redirect(http.StatusSeeOther, "/error")
		}
		return c.Redirect(http.StatusSeeOther, "/activate")
	}
	return c.Redirect(http.StatusSeeOther, "/login")
}
