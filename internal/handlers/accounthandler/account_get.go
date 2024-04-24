package accounthandler

import (
	"github.com/labstack/echo/v4"
	"github.com/rsvix/go-htmx-app-template/internal/templates"
	"gorm.io/gorm"
)

type getAccountHandlerParams struct {
	pageTitle string
	dbKey     string
}

func GetAccountHandler() *getAccountHandlerParams {
	return &getAccountHandlerParams{
		pageTitle: "Account",
		dbKey:     "__db",
	}
}

func (h getAccountHandlerParams) Serve(c echo.Context) error {
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

	return templates.AccountPage(c, h.pageTitle, result.Email, result.Firstname, result.Lastname).Render(c.Request().Context(), c.Response())
}
