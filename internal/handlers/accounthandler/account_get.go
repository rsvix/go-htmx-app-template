package accounthandler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rsvix/go-htmx-app-template/internal/templates"
	"github.com/rsvix/go-htmx-app-template/internal/utils"
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
	sessionInfo, err := utils.GetSessionInfo(c)
	if err != nil {
		log.Println(err)
		c.Response().Header().Set("HX-Redirect", "/error")
		return c.NoContent(http.StatusSeeOther)
	}

	db := c.Get(h.dbKey).(*gorm.DB)
	var result struct {
		Email     string
		Firstname string
		Lastname  string
	}
	// RAW
	_ = db.Raw("SELECT email, firstname, lastname FROM users WHERE id = ?;", sessionInfo.Id).Scan(&result)
	// GORM
	// _ = db.Table("users").Select("email", "firstname", "lastname").Where("id = ?", sessionInfo.Id).Scan(&result)

	return templates.AccountPage(c, h.pageTitle, result.Email, result.Firstname, result.Lastname).Render(c.Request().Context(), c.Response())
}
