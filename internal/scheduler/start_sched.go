package scheduler

import (
	"fmt"
	"log"
	"time"

	"github.com/go-co-op/gocron/v2"
	"gorm.io/gorm"
)

// https://github.com/go-co-op/gocron-gorm-lock

func BuildAsyncSched(db *gorm.DB) gocron.Scheduler {

	// Create cron lock table
	query := fmt.Sprintf(
		"CREATE TABLE IF NOT EXISTS cron_scheduler_lock (" +
			"id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT," +
			"instance_name VARCHAR(64) NOT NULL," +
			"last_ping VARCHAR(256)," +
			")")
	if err := db.Exec(query); err != nil {
		log.Println(err)
	}

	sched, err := gocron.NewScheduler()
	if err != nil {
		log.Fatal(err)
	}

	// add a job to the scheduler
	j, err := sched.NewJob(
		gocron.DurationJob(15*time.Second),
		gocron.NewTask(
			func(a string) {
				log.Println(a)
			},
			"\n\ntest\n\n",
		),
	)
	if err != nil {
		log.Println(err)
	}
	// each job has a unique id
	log.Println(j.ID())

	// start the scheduler
	sched.Start()

	return sched
}
