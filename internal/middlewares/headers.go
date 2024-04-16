package middlewares

import (
	"github.com/labstack/echo/v4"
)

func TextHTMLMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set("Content-Type", "text/html; charset=utf-8")
			return next(c)
		}
	}
}

func NoCache() echo.MiddlewareFunc {
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

// func ExecTime(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		before := time.Now()
// 		c.Response().Header().Set("ExecutionStartedAt", before.String())

// 		c.Response().Before(func() {
// 			after := time.Now()
// 			elapsed := time.Since(before)

// 			c.Response().Header().Set("ExecutionDoneAt", after.String())
// 			c.Response().Header().Set("ExecutionTime", elapsed.String())
// 		})

// 		if err := next(c); err != nil { //exec main process
// 			c.Error(err)
// 		}
// 		return nil
// 	}
// }
