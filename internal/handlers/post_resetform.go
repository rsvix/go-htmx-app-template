package handlers

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/rsvix/go-htmx-app-template/internal/hash"
	"github.com/rsvix/go-htmx-app-template/internal/templates"
	"gorm.io/gorm"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type postResetformHandlerParams struct {
}

func PostResetformHandler() *postResetformHandlerParams {
	return &postResetformHandlerParams{}
}

func (h *postResetformHandlerParams) Serve(c echo.Context) error {

	session, err := session.Get("authenticate-sessions", c)
	if err != nil {
		log.Printf("Get: %v\n", err)
	}

	if session.Values["pwreset"] != nil {
		if auth, ok := session.Values["pwreset"].(bool); !ok || !auth {
			return c.Redirect(http.StatusSeeOther, "/login")
		}
		password := c.Request().FormValue("password")
		password_conf := c.Request().FormValue("passwordconf")
		id := session.Values["user_id"].(string)

		log.Printf("id: %v\n", id)
		log.Printf("password: %v\n", password)
		log.Printf("password_conf: %v\n", password_conf)

		if password == password_conf {
			// hashPassword, err := hash.HashPasswordV1(password)
			hashPassword, err := hash.HashPasswordV2(password)
			if err != nil {
				log.Printf("Get: %v\n", err)
				return c.HTML(http.StatusInternalServerError, "<h2>Error, please try again</h2>")
			}
			log.Printf("hashPassword: %v\n", hashPassword)

			db := c.Get("__db").(*gorm.DB)
			// res := db.Table("users").Where("id = ?", id).Update("password", hashPassword)
			timein := time.Now().UTC().Add(1 * time.Hour)
			res := db.Table("users").Where("id = ?", id).Updates(map[string]interface{}{"password": hashPassword, "passwordchangetokenexpiration": timein})
			if res.Error != nil {
				return c.HTML(http.StatusUnprocessableEntity, "<h2>Error resseting password</h2>")
			}

			session.Values["pwreset"] = nil
			err = session.Save(c.Request(), c.Response())
			log.Printf("Save: %v\n", err)

			appName := os.Getenv("APP_NAME")
			return templates.MessagePage(c, appName, "Reset", "Your password was reset", true, false).Render(c.Request().Context(), c.Response())
		}
		return c.HTML(http.StatusUnprocessableEntity, "<h2>Passwords dont match</h2>")
	}
	return c.Redirect(http.StatusSeeOther, "/login")
}
