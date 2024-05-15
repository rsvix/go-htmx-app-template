package scheduledhandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rsvix/go-htmx-app-template/internal/structs"
	"github.com/rsvix/go-htmx-app-template/internal/templates"
	"github.com/rsvix/go-htmx-app-template/internal/utils"
	"gorm.io/gorm"
)

type getScheduledJobsHandlerParams struct {
	pageTitle string
}

func GetScheduledJobsHandler() *getScheduledJobsHandlerParams {
	return &getScheduledJobsHandlerParams{
		pageTitle: "CronJobs",
	}
}

func (h getScheduledJobsHandlerParams) Serve(c echo.Context) error {
	sessionInfo, err := utils.GetSessionInfo(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	db := c.Get("__db").(*gorm.DB)
	var result []structs.ScheduledJob
	db.Raw("SELECT * FROM scheduled_jobs;").Scan(&result)

	var m = make(map[string]bool)
	var agentArr []string
	for _, v := range result {
		agent := v.TargetAgent
		if !m[agent] {
			agentArr = append(agentArr, agent)
			m[agent] = true
		}
	}
	totalJobs := len(result)

	return templates.ScheduleJobsPage(c, h.pageTitle, sessionInfo.Id, sessionInfo.Username, result, agentArr, totalJobs).Render(c.Request().Context(), c.Response())
}
