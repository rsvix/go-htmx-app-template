package scheduler

import (
	"log"
	"time"

	"github.com/go-co-op/gocron/v2"
)

func BuildAsyncSched() gocron.Scheduler {

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
