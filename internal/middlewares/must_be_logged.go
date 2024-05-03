package middlewares

import (
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func MustBeLogged() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			session, err := session.Get("authenticate-sessions", c)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
				return c.Redirect(http.StatusSeeOther, "/login")
			}
			return next(c)
		}
	}
}
