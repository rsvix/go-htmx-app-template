package resethandler

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

type getResetTokenHandlerParams struct {
	queryParam string
	dbKey      string
}

func GetResetTokenHandler() *getResetTokenHandlerParams {
	return &getResetTokenHandlerParams{
		queryParam: "token",
		dbKey:      "__db",
	}
}

func (h getResetTokenHandlerParams) Serve(c echo.Context) error {
	// Path param - http://localhost:8080/pwreset/123qwert987012
	// token := c.Param(h.queryParam)

	// Query param - http://localhost:8080/pwreset?token=123qwert987012
	// token := c.QueryParam(h.queryParam)

	// Raw query - http://localhost:8080/pwreset?123qwert987012
	token := c.Request().URL.RawQuery

	if !strings.Contains(token, "O") {
		return c.Redirect(http.StatusSeeOther, "/login")
	}

	// Get session
	session, err := session.Get("authenticate-sessions", c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Check if session is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		extraInfo := strings.Split(token, "O")[1]

		decoded, err := hex.DecodeString(extraInfo)
		if err != nil {
			log.Println(err)
			return err
		}
		decodedStr := string(decoded[:])
		mode := strings.Split(decodedStr, "@")[0]
		id := strings.Split(decodedStr, "@")[1]

		if mode == "resetpwd" {
			session.Values["pwreset"] = false
			session.Values["pr_error"] = ""
			session.Values["user_id"] = ""

			db := c.Get("__db").(*gorm.DB)
			var result struct {
				Email                         string
				Passwordchangetoken           string
				Passwordchangetokenexpiration time.Time
				UserEnabled                   int
			}
			db.Raw("SELECT email, passwordchangetoken, passwordchangetokenexpiration, user_enabled FROM users WHERE id = ?", id).Scan(&result)

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

			if result.UserEnabled == 0 {
				return c.JSON(http.StatusBadRequest, "User must be activated first")
				// session.Values["pr_error"] = "User must be activated first"
				// err = session.Save(c.Request(), c.Response())
				// if err != nil {
				// 	log.Printf("Error saving session: %v\n", err)
				// }
				// return c.Redirect(http.StatusSeeOther, "/resetform")
			}

			if strings.Compare(strings.TrimSpace(result.Passwordchangetoken), strings.TrimSpace(token)) == 0 {
				session.Values["pwreset"] = true
				session.Values["user_id"] = id

				err = session.Save(c.Request(), c.Response())
				if err != nil {
					return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
				}
				return c.Redirect(http.StatusSeeOther, "/resetform")
			}
			return c.JSON(http.StatusBadRequest, "Invalid token")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "invalid 'mode' in token")
	}
	return c.Redirect(http.StatusSeeOther, "/")
}
