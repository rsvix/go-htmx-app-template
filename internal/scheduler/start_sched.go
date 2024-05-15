package scheduler

import (
	"fmt"
	"log"
	"time"

	"github.com/go-co-op/gocron/v2"
	"gorm.io/gorm"
)

// https://github.com/go-co-op/gocron-gorm-lock

func BuildAsyncSched(db *gorm.DB, instanceId string) gocron.Scheduler {

	// Create cron lock table
	query := fmt.Sprintf(
		"CREATE TABLE IF NOT EXISTS cron_scheduler_lock (" +
			"entry_id INT UNSIGNED NOT NULL," +
			"instance_id VARCHAR(64) NOT NULL," +
			"last_update TIMESTAMP NOT NULL," +
			"UNIQUE (entry_id)" +
			")")
	if err := db.Exec(query); err.Error != nil {
		log.Panicf("Error creating cron_scheduler_lock table: %v", err)
	}

	t := time.Now().UTC().Format("2006-01-02 15:04:05")
	log.Println(t)
	result := db.Raw("INSERT INTO cron_scheduler_lock (entry_id, instance_id, last_update) VALUES (?, ?, ?);", 1, instanceId, t)
	log.Printf("RowsAffected: %v", result.RowsAffected)
	log.Printf("Error: %v", result.Error)

	var schedInfo []struct {
		EntryId    uint
		InstanceId string
		LastUpdate time.Time
	}
	db.Raw("SELECT * FROM cron_scheduler_lock;").Scan(&schedInfo)
	log.Printf("schedInfo: %v", schedInfo)

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
	// 				LastUpdate   string
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
