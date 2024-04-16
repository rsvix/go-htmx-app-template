package handlers

import (
	"log"
	"net/http"
	"net/mail"
	"strconv"

	"github.com/rsvix/go-htmx-app-template/internal/hash"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type postLoginHandlerParams struct {
	dbKey string
}

func PostLoginHandler() *postLoginHandlerParams {
	return &postLoginHandlerParams{
		dbKey: "__db",
	}
}

func (h *postLoginHandlerParams) Serve(c echo.Context) error {

	email := c.Request().FormValue("email")
	password := c.Request().FormValue("password")
	remember := c.Request().FormValue("remember")
	log.Printf("remember: %v\n", remember)

	// https://stackoverflow.com/questions/2185951/how-do-i-keep-a-user-logged-into-my-site-for-months

	if _, err := mail.ParseAddress(email); err != nil {
		return c.HTML(http.StatusUnprocessableEntity, "<h2>Invalid credentials</h2>")
	}

	db := c.Get(h.dbKey).(*gorm.DB)
	var result struct {
		Id        int
		Firstname string
		Password  string
		Enabled   int
	}
	// db.Raw("SELECT password, enabled FROM users WHERE email = $1", email).Scan(&result)
	db.Table("users").Select("id", "firstname", "password", "enabled").Where("email = ?", email).Scan(&result)
	// log.Printf("result: %v\n", result)

	if result.Id != 0 {
		if result.Enabled == 0 {
			return c.HTML(http.StatusUnprocessableEntity, "<h2>User not enabled<br/>Check your email</h2>")
		}

		// if hash.CheckPasswordHashV1(password, result.Password) {
		if hash.CheckPasswordHashV2(password, result.Password) {

			// Get session
			session, err := session.Get("authenticate-sessions", c)
			if err != nil {
				log.Printf("Error getting session: %v\n", err)
				return c.HTML(http.StatusInternalServerError, "<h2>Error, please try again</h2>")
			}

			// Set session values
			session.Values["authenticated"] = true
			session.Values["email"] = email
			session.Values["id"] = strconv.FormatUint(uint64(result.Id), 10)
			session.Values["firstname"] = result.Firstname

			// Save updated session
			if err := session.Save(c.Request(), c.Response()); err != nil {
				log.Printf("Error saving session: %s", err)
				return c.HTML(http.StatusInternalServerError, "<h2>Error, please try again</h2>")
			}

			// return c.Redirect(http.StatusSeeOther, "/")
			c.Response().Header().Set("HX-Redirect", "/")
			return c.NoContent(http.StatusOK)
		}
		return c.HTML(http.StatusUnprocessableEntity, "<h2>Invalid credentials</h2>")
	}
	return c.HTML(http.StatusUnprocessableEntity, "<h2>Invalid credentials</h2>")
}
