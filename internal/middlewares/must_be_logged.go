package middlewares

import (
	"log"
	"os"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/rsvix/go-htmx-app-template/internal/templates"
)

type loginPageParams struct {
	appName   string
	pageTitle string
}

func LoginPageParamsHandler() *loginPageParams {
	return &loginPageParams{
		appName:   os.Getenv("APP_NAME"),
		pageTitle: "Login",
	}
}

func (h loginPageParams) MustBeLogged() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			session, err := session.Get("authenticate-sessions", c)
			if err != nil {
				log.Printf("Error getting session: %v\n", err)
				return err
			}

			if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
				return templates.LoginPage(h.appName, h.pageTitle).Render(c.Request().Context(), c.Response())
			}

			return next(c)
		}
	}
}
