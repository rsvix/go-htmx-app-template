package resethandler

import (
	"fmt"
	"log"
	"net/http"
	"net/mail"
	"os"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rsvix/go-htmx-app-template/internal/emails"
	"github.com/rsvix/go-htmx-app-template/internal/hash"
	"gorm.io/gorm"
)

type postResetHandlerParams struct {
}

func PostResetHandler() *postResetHandlerParams {
	return &postResetHandlerParams{}
}

func (h *postResetHandlerParams) Serve(c echo.Context) error {

	email := c.Request().FormValue("email")
	_, err := mail.ParseAddress(email)
	if err != nil {
		return c.HTML(http.StatusUnprocessableEntity, "<h2>Couldn't process your request</h2>")
	}
	log.Printf("email: %s\n", email)

	db := c.Get("__db").(*gorm.DB)
	var result struct {
		Id                            int
		Passwordchangetokenexpiration time.Time
	}
	db.Raw("SELECT id, passwordchangetokenexpiration FROM users WHERE email = ?", email).Scan(&result)

	if result.Id != 0 {
		id := strconv.FormatUint(uint64(result.Id), 10)

		diff := result.Passwordchangetokenexpiration.Sub(time.Now().UTC())
		// diff := time.Now().UTC().Sub(result.Passwordchangetokenexpiration)
		secs := diff.Seconds()
		log.Printf("diff: %v - secs: %v\n", diff, secs)
		if secs > 0.0 {
			return c.HTML(http.StatusInternalServerError, "<h2>Can't request multiple password<br/>resets in a short period</h2>")
		}

		// resetToken, err := hash.GenerateActivationToken()
		resetToken, err := hash.GenerateToken(false, id)
		if err != nil {
			log.Println(err)
			return c.HTML(http.StatusInternalServerError, "<h2>Couldn't process your request</h2>")
		}
		log.Printf("passToken: %v\n", resetToken)
		timein := time.Now().UTC().Add(1 * time.Hour)

		result := db.Exec("UPDATE users SET passwordchangetoken = ?, passwordchangetokenexpiration = ? WHERE id = ?;", resetToken, timein, id)
		log.Printf("Query result: %v\n", result)

		var dbEmail string
		db.Raw("SELECT email from users WHERE id = ?;", id).Scan(&dbEmail)
		log.Printf("dbEmail: %v\n", dbEmail)

		appPort, _ := os.LookupEnv("APP_PORT")
		passUrl := fmt.Sprintf("http://localhost:%s/pwreset?%s", appPort, resetToken)
		log.Printf("passUrl: %v\n", passUrl)

		// Must configure SMTP server or other email sending service
		if _, ok := os.LookupEnv("SENDER_PSWD"); ok {
			if err := emails.SendResetMail(email, passUrl, emails.DefaultParams()); err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			return c.HTML(http.StatusOK, fmt.Sprintf("<h2>Email sent to %s, id: %s</h2>", email, id))
		}

		// For testing without email sender configured
		msg := fmt.Sprintf("<h2><div style=\"text-decoration-line: underline;\"><a href=\"%s\">Reset your password<br/>by clicking here</a></div></h2>", passUrl)
		return c.HTML(http.StatusOK, msg)
	}
	return c.HTML(http.StatusOK, "<h2>Couldn't process your request</h2>")
}
