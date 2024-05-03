package indexhandler

import (
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
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return templates.IndexPage(c, h.pageTitle, sessionInfo.Username).Render(c.Request().Context(), c.Response())
}
