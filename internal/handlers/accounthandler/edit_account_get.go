package accounthandler

import (
	"github.com/labstack/echo/v4"
	"github.com/rsvix/go-htmx-app-template/internal/templates"
	"gorm.io/gorm"
)

type getEditAccountHandlerParams struct {
	pageTitle string
	dbKey     string
}

func GetEditAccountHandler() *getEditAccountHandlerParams {
	return &getEditAccountHandlerParams{
		pageTitle: "Account",
		dbKey:     "__db",
	}
}

func (h getEditAccountHandlerParams) Serve(c echo.Context) error {
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
