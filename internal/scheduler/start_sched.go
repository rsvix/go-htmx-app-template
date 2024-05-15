package scheduler

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-co-op/gocron/v2"
	"gorm.io/gorm"
)

// https://github.com/go-co-op/gocron-gorm-lock

func BuildAsyncSched(db *gorm.DB, instanceId string) gocron.Scheduler {

	// Create cron lock table
	query := fmt.Sprintf(
		"CREATE TABLE IF NOT EXISTS cron_scheduler_lock (" +
			"id INT UNSIGNED PRIMARY KEY NOT NULL," +
			"instance_id VARCHAR(64) NOT NULL," +
			"last_update TIMESTAMP NOT NULL," +
			"UNIQUE (id)" +
			")")
	if err := db.Exec(query); err.Error != nil {
		log.Panicf("Error creating cron_scheduler_lock table: %v", err)
	}

	t := time.Now().UTC()
	result := db.Exec("INSERT IGNORE INTO cron_scheduler_lock (id, instance_id, last_update) VALUES ('1', ?, ?);", instanceId, t)
	if result.RowsAffected == 1 {
		os.Setenv("IS_SCHEDULER", "true")
	} else {
		os.Setenv("IS_SCHEDULER", "false")
	}

	sched, err := gocron.NewScheduler()
	if err != nil {
		log.Fatal(err)
	}

	// Job that defines the instance responsible for creating tasks
	j, err := sched.NewJob(
		gocron.DurationJob(15*time.Second),
		gocron.NewTask(
			func() {
				var schedInfo struct {
					Id         uint
					InstanceId string
					LastUpdate time.Time
				}
				db.Raw("SELECT * FROM cron_scheduler_lock WHERE id = '1';").Scan(&schedInfo)
				log.Printf("schedInfo: %v", schedInfo)
				if schedInfo.InstanceId == instanceId {
					t := time.Now().UTC()
					db.Exec("UPDATE cron_scheduler_lock SET last_update = ? WHERE id = '1';", t)
				} else {
					os.Setenv("IS_SCHEDULER", "false")
					diffInSecs := time.Now().UTC().Sub(schedInfo.LastUpdate).Seconds()
					if diffInSecs > 60 {
						log.Printf("diffInSecs: %v", diffInSecs)
						db.Exec("UPDATE cron_scheduler_lock SET instance_id = ?, last_update = ? WHERE id = '1';", instanceId, t)
						os.Setenv("IS_SCHEDULER", "true")
					}
				}
			},
		),
	)
	if err != nil {
		log.Println(err)
	}
	log.Printf("scheduler_lock job id: %v", j.ID())

	// Test job
	jt, err := sched.NewJob(
		gocron.DurationJob(15*time.Second),
		gocron.NewTask(
			func(a string) {
				if os.Getenv("IS_SCHEDULER") == "true" {
					log.Println(a)
				}
			},
			" ##########################  test  ##########################",
		),
	)
	if err != nil {
		log.Println(err)
	}
	log.Printf("test job id: %v", jt.ID())

	var schedJobsTable []struct {
		Id          uint
		CronExp     string
		CronDesc    string
		BotName     string
		BotVersion  string
		TargetAgent string
		Params      string
		Uuid        string
	}
	db.Raw("SELECT * FROM scheduled_jobs;").Scan(&schedJobsTable)
	log.Printf("schedJobsTable: %v", schedJobsTable)

	for _, job := range schedJobsTable {
		log.Printf("job: %v", job)
	}

	// start the scheduler
	sched.Start()
	return sched
}
