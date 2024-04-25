package resethandler

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

type getTokenHandlerParams struct {
	queryParam string
	dbKey      string
}

func GetTokenHandler() *getTokenHandlerParams {
	return &getTokenHandlerParams{
		queryParam: "token",
		dbKey:      os.Getenv("DB_CONTEXT_KEY"),
	}
}

func (h getTokenHandlerParams) Serve(c echo.Context) error {
	token := c.Param(h.queryParam)

	if !strings.Contains(token, "O") {
		return c.Redirect(http.StatusSeeOther, "/login")
	}

	session, err := session.Get("authenticate-sessions", c)
	if err != nil {
		log.Printf("Error getting session: %v\n", err)
	}

	if auth, ok := session.Values["authenticated"].(bool); ok && !auth {
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

			db := c.Get(h.dbKey).(*gorm.DB)

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

			log.Printf("token: %v\n", token)
			log.Printf("Passwordchangetoken: %v\n", result.Passwordchangetoken)

			if strings.Compare(strings.TrimSpace(result.Passwordchangetoken), strings.TrimSpace(token)) == 0 {
				session.Values["pwreset"] = true
				session.Values["user_id"] = id
				err = session.Save(c.Request(), c.Response())
				if err != nil {
					log.Printf("Error saving session: %v\n", err)
				}
				return c.Redirect(http.StatusSeeOther, "/resetform")
			}
			return c.JSON(http.StatusBadRequest, "Invalid token")
		}
		return c.Redirect(http.StatusSeeOther, "/error")
	}
	return c.Redirect(http.StatusSeeOther, "/")
}
