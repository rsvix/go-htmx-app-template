package handlers

import (
	"log"
	"net/http"
	"net/mail"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/rsvix/go-htmx-app-template/internal/hash"
	"github.com/rsvix/go-htmx-app-template/internal/utils"
	"gorm.io/gorm"
)

type postEditAccountHandlerParams struct {
	pageTitle string
	dbKey     string
}

func PostEditAccountHandler() *postEditAccountHandlerParams {
	return &postEditAccountHandlerParams{
		pageTitle: "Account",
		dbKey:     "__db",
	}
}

func (h postEditAccountHandlerParams) Serve(c echo.Context) error {

	email := c.Request().FormValue("email")
	firstname := c.Request().FormValue("firstname")
	lastname := c.Request().FormValue("lastname")
	password := c.Request().Header.Get("HX-Prompt")

	log.Printf("\nemail: %s\nfirstname: %s\nlastname: %s\npassword: %s\n", email, firstname, lastname, password)

	if _, err := mail.ParseAddress(email); err != nil {
		return c.HTML(http.StatusUnprocessableEntity, "<p>Invalid email</p>")
	}

	if !utils.IsValidName(firstname) {
		return c.HTML(http.StatusUnprocessableEntity, "<p>Invalid first name</p>")
	}

	if !utils.IsValidName(lastname) {
		return c.HTML(http.StatusUnprocessableEntity, "<p>Invalid lastname</p>")
	}

	session, err := session.Get("authenticate-sessions", c)
	if err != nil {
		log.Printf("Error getting session: %v\n", err)
		return err
	}

	var id string
	if value, ok := session.Values["id"].(string); ok {
		id = value
	} else {
		return err
	}

	db := c.Get(h.dbKey).(*gorm.DB)
	var result struct {
		Password string
	}
	db.Table("users").Select("password").Where("id = ?", email).Scan(&result)

	if hash.CheckPasswordHashV2(password, result.Password) {
		res := db.Table("users").Where("id = ?", id).Updates(map[string]interface{}{"email": email, "firstname": firstname, "lastname": lastname})
		if res.Error != nil {
			return c.HTML(http.StatusUnprocessableEntity, "<p>Error updating account</p>")
		}
		return c.Redirect(http.StatusSeeOther, "/account")
	}
	return c.HTML(http.StatusUnprocessableEntity, "<p>Invalid password</p>")
}
