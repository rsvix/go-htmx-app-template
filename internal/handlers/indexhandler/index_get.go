package indexhandler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rsvix/go-htmx-app-template/internal/templates"
	"github.com/rsvix/go-htmx-app-template/internal/utils"
)

type getIndexHandlerParams struct {
	pageTitle string
}

func GetIndexHandler() *getIndexHandlerParams {
	return &getIndexHandlerParams{
		pageTitle: "Index",
	}
}

func (h getIndexHandlerParams) Serve(c echo.Context) error {
	sessionInfo, err := utils.GetSessionInfo(c)
	if err != nil {
		log.Printf("Error getting session info: %v\n", err)
		c.Response().Header().Set("HX-Redirect", "/error")
		return c.NoContent(http.StatusSeeOther)
	}

	return templates.IndexPage(c, h.pageTitle, sessionInfo.Username).Render(c.Request().Context(), c.Response())
}
