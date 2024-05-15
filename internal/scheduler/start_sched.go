package scheduler

import (
	"fmt"
	"log"
	"time"

	"github.com/go-co-op/gocron/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// https://github.com/go-co-op/gocron-gorm-lock

func BuildAsyncSched(db *gorm.DB, instanceId uuid.UUID) gocron.Scheduler {

	// Create cron lock table
	query := fmt.Sprintf(
		"CREATE TABLE IF NOT EXISTS cron_scheduler_lock (" +
			"id INT UNSIGNED PRIMARY KEY NOT NULL," +
			"instance_id VARCHAR(64) NOT NULL," +
			"last_ping DATETIME," +
			"UNIQUE (id)" +
			")")
	if err := db.Exec(query); err.Error != nil {
		log.Panicf("Error creating cron_scheduler_lock table: %v", err)
	}

	t := time.Now().UTC()
	var instanceIdInTable string
	db.Raw("INSERT IGNORE INTO cron_scheduler_lock (id, instance_id, last_ping) VALUES (?, ?, ?) RETURNING instance_id;", 1, instanceId, t).Scan(&instanceIdInTable)
	log.Printf("instanceIdInTable: %v", instanceIdInTable)

	sched, err := gocron.NewScheduler()
	if err != nil {
		log.Fatal(err)
	}

	// j, err := sched.NewJob(
	// 	gocron.DurationJob(10*time.Second),
	// 	gocron.NewTask(
	// 		func() {
	// 			var schedInfo []struct {
	// 				Id         uint
	// 				InstanceId string
	// 				LastPing   string
	// 			}
	// 			db.Raw("SELECT * FROM cron_scheduler_lock;").Scan(&schedInfo)
	// 		},
	// 	),
	// )
	// if err != nil {
	// 	log.Println(err)
	// }
	// // each job has a unique id
	// log.Println(j.ID())

	// add a job to the scheduler
	jt, err := sched.NewJob(
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
	log.Println(jt.ID())

	// start the scheduler
	sched.Start()

	return sched
}
