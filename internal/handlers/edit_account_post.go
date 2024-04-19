package handlers

import (
	"log"
	"net/http"
	"net/mail"

	"github.com/labstack/echo/v4"
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

	log.Printf("email: %s\nfirstname: %s\nlastname: %s\n", email, firstname, lastname)

	if _, err := mail.ParseAddress(email); err != nil {
		return c.HTML(http.StatusUnprocessableEntity, "<p>Invalid email</p>")
	}

	if !utils.IsValidName(firstname) {
		return c.HTML(http.StatusUnprocessableEntity, "<p>Invalid first name</p>")
	}

	if !utils.IsValidName(lastname) {
		return c.HTML(http.StatusUnprocessableEntity, "<p>Invalid lastname</p>")
	}

	var id string
	if value, ok := c.Get("userId").(string); ok {
		id = value
	}

	db := c.Get(h.dbKey).(*gorm.DB)
	res := db.Table("users").Where("id = ?", id).Updates(map[string]interface{}{"email": email, "firstname": firstname, "lastname": lastname})
	if res.Error != nil {
		return c.HTML(http.StatusUnprocessableEntity, "<p>Error updating account</p>")
	}

	return c.Redirect(http.StatusSeeOther, "/account")
}
