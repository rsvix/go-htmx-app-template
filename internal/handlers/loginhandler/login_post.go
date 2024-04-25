package loginhandler

import (
	"log"
	"net/http"
	"net/mail"

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

	if _, err := mail.ParseAddress(email); err != nil {
		return c.HTML(http.StatusUnprocessableEntity, "Invalid credentials")
	}

	db := c.Get(h.dbKey).(*gorm.DB)
	var user struct {
		Id       int
		Username string
		Password string
		Enabled  int
	}
	result := db.Raw("SELECT id, username, password, enabled FROM users WHERE email = ?;", email).Scan(&user)
	log.Printf("result: %v\n", result)

	if user.Id != 0 {
		if user.Enabled == 0 {
			return c.HTML(http.StatusUnprocessableEntity, "User not enabled<br/>Check your email")
		}

		// if hash.CheckPasswordHashV1(password, result.Password) {
		if hash.CheckPasswordHashV2(password, user.Password) {

			session, err := session.Get("authenticate-sessions", c)
			if err != nil {
				log.Printf("Error getting session: %v\n", err)
				return c.HTML(http.StatusInternalServerError, "Error, please try again")
			}

			session.Values["authenticated"] = true
			session.Values["user_id"] = user.Id
			// session.Values["user_id"] = strconv.FormatUint(uint64(user.Id), 10)
			session.Values["user_email"] = email
			session.Values["username"] = user.Username

			if remember == "true" {
				session.Options.MaxAge = 84600 * 30
			}

			if err := session.Save(c.Request(), c.Response()); err != nil {
				log.Printf("Error saving session: %s", err)
				return c.HTML(http.StatusInternalServerError, "Error, please try again")
			}

			c.Response().Header().Set("HX-Redirect", "/")
			return c.NoContent(http.StatusSeeOther)
		}
		return c.HTML(http.StatusUnprocessableEntity, "Invalid credentials")
	}
	return c.HTML(http.StatusUnprocessableEntity, "Invalid credentials")
}
