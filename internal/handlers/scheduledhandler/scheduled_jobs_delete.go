package scheduledhandler

import (
	"log"
	"net/http"

	"github.com/go-co-op/gocron/v2"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type deleteScheduledJobsHandlerParams struct {
	pageTitle string
	sched     gocron.Scheduler
}

func DeleteScheduledJobsHandler(sched gocron.Scheduler) *deleteScheduledJobsHandlerParams {
	return &deleteScheduledJobsHandlerParams{
		pageTitle: "CronJobs",
		sched:     sched,
	}
}

func (h deleteScheduledJobsHandlerParams) Serve(c echo.Context) error {
	jobId := c.Param("id")
	db := c.Get("__db").(*gorm.DB)

	var jobUuidStr string
	db.Raw("SELECT uuid FROM scheduled_jobs WHERE id = ?;", jobId).Scan(&jobUuidStr)

	jobUuid, _ := uuid.Parse(jobUuidStr)
	if err := h.sched.RemoveJob(jobUuid); err != nil {
		log.Printf("Error removing cronjob with id '%v': %v", jobId, err)
	}
	db.Exec("DELETE FROM scheduled_jobs WHERE id = ?", jobId)

	c.Response().Header().Set("HX-Redirect", "/cronjobs")
	return c.NoContent(http.StatusSeeOther)
}
