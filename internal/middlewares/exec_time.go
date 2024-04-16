package middlewares

import (
	"time"

	"github.com/labstack/echo/v4"
)

func ExecTime(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		before := time.Now()
		c.Response().Header().Set("ExecutionStartedAt", before.String())

		c.Response().Before(func() {
			after := time.Now()
			elapsed := time.Since(before)

			c.Response().Header().Set("ExecutionDoneAt", after.String())
			c.Response().Header().Set("ExecutionTime", elapsed.String())
		})

		if err := next(c); err != nil { //exec main process
			c.Error(err)
		}
		return nil
	}
}
