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
	"github.com/rsvix/go-htmx-app-template/internal/templates"
	"gorm.io/gorm"
)

type getActivationTokenHandlerParams struct {
	queryParam string
	dbKey      string
	pageTitle  string
}

func GetActivationTokenHandler() *getActivationTokenHandlerParams {
	return &getActivationTokenHandlerParams{
		queryParam: "token",
		dbKey:      os.Getenv("DB_CONTEXT_KEY"),
		pageTitle:  "Activate",
	}
}

func (h getActivationTokenHandlerParams) Serve(c echo.Context) error {
	// Path param - http://localhost:8080/activation/123qwert987012
	// token := c.Param(h.queryParam)

	// Query param - http://localhost:8080/activation?token=123qwert987012
	// token := c.QueryParam(h.queryParam)

	// Raw query - http://localhost:8080/activation?123qwert987012
	token := c.Request().URL.RawQuery

	if !strings.Contains(token, "O") || token == "" {
		return c.Redirect(http.StatusSeeOther, "/login")
	}

	extraInfo := strings.Split(token, "O")[1]
	decoded, err := hex.DecodeString(extraInfo)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	decodedStr := string(decoded[:])
	mode := strings.Split(decodedStr, "@")[0]
	id := strings.Split(decodedStr, "@")[1]

	if mode == "activate" {
		db := c.Get(h.dbKey).(*gorm.DB)
		var result struct {
			Activationtoken           string
			Activationtokenexpiration time.Time
			Enabled                   int
		}
		_ = db.Raw("SELECT activationtoken, activationtokenexpiration, user_enabled FROM users WHERE id = ?", id).Scan(&result)

		if result.Enabled == 0 {
			diff := result.Activationtokenexpiration.Sub(time.Now().UTC())
			secs := diff.Seconds()
			// log.Printf("diff: %v\nsecs: %v\n", diff, secs)

			if secs > 0.0 {
				if strings.Compare(strings.TrimSpace(result.Activationtoken), strings.TrimSpace(token)) == 0 {
					timeNow := time.Now().UTC()
					res := db.Table("users").Where("id = ?", id).Updates(map[string]interface{}{"user_enabled": 1, "activationtokenexpiration": timeNow})
					if res.Error != nil {
						log.Println("Error enabling user")
						return c.Redirect(http.StatusSeeOther, "/error")
					}
					return templates.ActivationPage(c, h.pageTitle, true, "Account activated").Render(c.Request().Context(), c.Response())
				} else {
					return templates.ActivationPage(c, h.pageTitle, false, "Invalid token").Render(c.Request().Context(), c.Response())
				}
			} else {
				session, err := session.Get("authenticate-sessions", c)
				if err != nil {
					log.Printf("Error getting session: %v\n", err)
				}
				session.Values["user_id"] = id
				if err := session.Save(c.Request(), c.Response()); err != nil {
					return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
				}
				return templates.ActivationPage(c, h.pageTitle, false, "Token expired").Render(c.Request().Context(), c.Response())
			}
		} else {
			return c.Redirect(http.StatusSeeOther, "/login")
		}
	}
	return c.Redirect(http.StatusSeeOther, "/login")
}
