package accounthandler

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/rsvix/go-htmx-app-template/internal/templates"
	"github.com/rsvix/go-htmx-app-template/internal/utils"
	"gorm.io/gorm"
)

type getEditAccountHandlerParams struct {
	pageTitle string
	dbKey     string
}

func GetEditAccountHandler() *getEditAccountHandlerParams {
	return &getEditAccountHandlerParams{
		pageTitle: "Account",
		dbKey:     os.Getenv("DB_CONTEXT_KEY"),
	}
}

func (h getEditAccountHandlerParams) Serve(c echo.Context) error {
	sessionInfo, err := utils.GetSessionInfo(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
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

	return templates.EditAccountPage(c, h.pageTitle, result.Email, result.Firstname, result.Lastname).Render(c.Request().Context(), c.Response())
}
