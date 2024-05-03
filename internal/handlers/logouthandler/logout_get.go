package logouthandler

import (
	"net/http"
	"time"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type getLogoutHandlerParams struct {
}

func GetLogoutHandler() *getLogoutHandlerParams {
	return &getLogoutHandlerParams{}
}

func (h getLogoutHandlerParams) Serve(c echo.Context) error {
	c.Response().Header().Set("Cache-Control", "no-cache, private, max-age=0")
	c.Response().Header().Set("Expires", time.Unix(0, 0).Format(http.TimeFormat))
	c.Response().Header().Set("Pragma", "no-cache")
	c.Response().Header().Set("X-Accel-Expires", "0")

	session, err := session.Get("authenticate-sessions", c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	session.Values["authenticated"] = false
	session.Values["user_id"] = nil
	session.Values["user_email"] = nil
	session.Values["username"] = nil
	session.Options.MaxAge = -1

	if err := session.Save(c.Request(), c.Response()); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.Redirect(http.StatusSeeOther, "/login")
}
