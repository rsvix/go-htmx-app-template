package middlewares

import (
	"github.com/labstack/echo/v4"
)

// https://github.com/TomDoesTech/GOTTH/blob/main/internal/middleware/middleware.go

func TextHTMLMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set("Content-Type", "text/html; charset=utf-8")
			return next(c)
		}
	}
}
