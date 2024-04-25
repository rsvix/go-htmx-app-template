package accounthandler

import (
	"net/http"
	"net/mail"
	"os"

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
		dbKey:     os.Getenv("DB_CONTEXT_KEY"),
	}
}

func (h postEditAccountHandlerParams) Serve(c echo.Context) error {

	email := c.Request().FormValue("email")
	firstname := c.Request().FormValue("firstname")
	lastname := c.Request().FormValue("lastname")
	password := c.Request().Header.Get("HX-Prompt")

	if _, err := mail.ParseAddress(email); err != nil {
		return c.HTML(http.StatusUnprocessableEntity, "Invalid email")
	}

	if !utils.IsValidName(firstname) {
		return c.HTML(http.StatusUnprocessableEntity, "Invalid first name")
	}

	if !utils.IsValidName(lastname) {
		return c.HTML(http.StatusUnprocessableEntity, "Invalid lastname")
	}

	sessionInfo, err := utils.GetSessionInfo(c)
	if err != nil {
		c.Response().Header().Set("HX-Redirect", "/error")
		return c.NoContent(http.StatusSeeOther)
	}

	db := c.Get(h.dbKey).(*gorm.DB)
	var result struct {
		Password string
	}
	db.Table("users").Select("password").Where("id = ?", sessionInfo.Id).Scan(&result)

	if hash.CheckPasswordHashV2(password, result.Password) {
		res := db.Table("users").Where("id = ?", sessionInfo.Id).Updates(map[string]interface{}{"email": email, "firstname": firstname, "lastname": lastname})
		if res.Error != nil {
			return c.HTML(http.StatusUnprocessableEntity, "Error updating account")
		}
		return c.Redirect(http.StatusSeeOther, "/account")
	}
	return c.HTML(http.StatusUnprocessableEntity, "Invalid password")
}
