package scheduledhandler

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-co-op/gocron/v2"
	"github.com/labstack/echo/v4"
	"github.com/rsvix/go-htmx-app-template/internal/crondescriptor"
	"gorm.io/gorm"
)

type postNewJobHandlerParams struct {
	pageTitle string
	sched     gocron.Scheduler
}

func PostNewJobHandler(sched gocron.Scheduler) *postNewJobHandlerParams {
	return &postNewJobHandlerParams{
		pageTitle: "CronJobs",
		sched:     sched,
	}
}

func (h postNewJobHandlerParams) Serve(c echo.Context) error {
	// sessionInfo, err := utils.GetSessionInfo(c)
	// if err != nil {
	// 	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	// }

	cronExp := c.Request().FormValue("cronExp")
	botName := c.Request().FormValue("botName")
	botVersion := c.Request().FormValue("botVersion")
	agentName := c.Request().FormValue("agentName")

	db := c.Get("__db").(*gorm.DB)

	if len(strings.Split(cronExp, " ")) != 5 || len(strings.Split(cronExp, " ")) != 6 {
		log.Println("Cron expression must have 5 or 6 fields")
		log.Println(len(strings.Split(cronExp, " ")))
		c.Response().Header().Set("HX-Redirect", "/cronjobs")
		return c.NoContent(http.StatusSeeOther)
	}

	cd, err := crondescriptor.NewCronDescriptor(cronExp)
	if err != nil {
		log.Printf("error creating descriptor: %v\n", err.Error())
		c.Response().Header().Set("HX-Redirect", "/cronjobs")
		return c.NoContent(http.StatusSeeOther)
	}

	fullDescription, err := cd.GetDescription(crondescriptor.Full)
	if err != nil {
		log.Printf("error getting description: %v\n", err.Error())
		c.Response().Header().Set("HX-Redirect", "/cronjobs")
		return c.NoContent(http.StatusSeeOther)
	}
	fmt.Printf("%s => %s\n", cronExp, *fullDescription)

	job, err := h.sched.NewJob(
		gocron.CronJob(cronExp, true),
		gocron.NewTask(
			func() {
				if os.Getenv("IS_SCHEDULER") == "true" {
					cmd := fmt.Sprintf("exec-bot::%v::%v", botName, botVersion)
					log.Printf("cmd: %v", cmd)
				}
			},
		),
	)
	if err != nil {
		log.Println(err)
	}

	result := db.Exec("INSERT INTO scheduled_jobs (cron_exp, cron_desc, bot_name, bot_version, target_agent, uuid) VALUES (?, ?, ?, ?, ?, ?);",
		cronExp,
		&fullDescription,
		botName,
		botVersion,
		agentName,
		job.ID().String(),
	)
	if err := result.Error; err != nil {
		log.Printf("error updating db: %v\n", err.Error())
		return c.Redirect(http.StatusSeeOther, "/error")
	}
	c.Response().Header().Set("HX-Redirect", "/cronjobs")
	return c.NoContent(http.StatusSeeOther)
}
