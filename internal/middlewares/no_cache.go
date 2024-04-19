package middlewares

import "github.com/labstack/echo/v4"

func NoCacheHeaders() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set("Cache-Control", "no-cache, private, max-age=0")
			// c.Response().Header().Set("Expires", time.Unix(0, 0).Format(http.TimeFormat))
			c.Response().Header().Set("Pragma", "no-cache")
			c.Response().Header().Set("X-Accel-Expires", "0")
			return next(c)
		}
	}
}
