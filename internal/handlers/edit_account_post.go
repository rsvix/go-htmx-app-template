package handlers

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/rsvix/go-htmx-app-template/internal/templates"
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

	var id string
	if value, ok := c.Get("userId").(string); ok {
		id = value
	}
	db := c.Get(h.dbKey).(*gorm.DB)
	var result struct {
		Email     string
		Firstname string
		Lastname  string
	}
	db.Table("users").Select("email", "firstname", "lastname").Where("id = ?", id).Scan(&result)

	return templates.EditAccountPage(c, h.pageTitle, result.Email, result.Firstname, result.Lastname).Render(c.Request().Context(), c.Response())
}
