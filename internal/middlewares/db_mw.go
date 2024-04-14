package middlewares

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

const (
	dbContextKey = "__db"
)

func DatabaseMiddleware(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(dbContextKey, db)
			return next(c)
		}
	}
}
