package handlers

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

type PostResetHandler struct {
}

type PostResetHandlerParams struct {
}

func NewPostResetHandler(params PostResetHandlerParams) *PostResetHandler {
	return &PostResetHandler{}
}

func (h *PostResetHandler) ServeHTTP(c echo.Context) error {

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

		resetToken, err := hash.GenerateActivationToken()
		if err != nil {
			log.Println(err)
			return c.HTML(http.StatusInternalServerError, "<h2>Couldn't process your request</h2>")
		}
		log.Printf("passToken: %v\n", resetToken)
		timein := time.Now().UTC().Add(1 * time.Hour)

		var dbEmail string
		result := db.Raw("UPDATE users SET passwordchangetoken = ?, passwordchangetokenexpiration = ? WHERE id = ? RETURNING email;", resetToken, timein, id).Scan(&dbEmail)
		log.Printf("dbEmail: %v\n", dbEmail)
		log.Printf("Query result: %v\n", result)

		appPort, _ := os.LookupEnv("APP_PORT")
		passUrl := fmt.Sprintf("http://localhost:%s/pwreset/%s/%s", appPort, id, resetToken)
		log.Printf("passUrl: %v\n", passUrl)

		// Must configure SMTP server or other email sending service
		if _, ok := os.LookupEnv("SENDER_PSWD"); ok {
			emails.SendResetEmail(email, passUrl)
			return c.HTML(http.StatusOK, fmt.Sprintf("<h2>Email sent to %s, id: %s</h2>", email, id))
		}

		// For testing without email sender configured
		msg := fmt.Sprintf("<h2><div style=\"text-decoration-line: underline;\"><a href=\"%s\">Reset your password<br/>by clicking here</a></div></h2>", passUrl)
		return c.HTML(http.StatusOK, msg)
	}
	return c.HTML(http.StatusOK, "<h2>Couldn't process your request</h2>")
}
