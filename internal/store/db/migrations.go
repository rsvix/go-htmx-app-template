package db

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

func MigrateDb(db *gorm.DB) {

	createUsersTable := fmt.Sprintf(
		"CREATE TABLE IF NOT EXISTS users (" +
			"id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT," +
			"email VARCHAR(64) NOT NULL," +
			"username VARCHAR(64) NOT NULL," +
			"firstname VARCHAR(64) NOT NULL," +
			"lastname VARCHAR(64) NOT NULL," +
			"password VARCHAR(128) NOT NULL," +
			"activationtoken VARCHAR(256)," +
			"activationtokenexpiration TIMESTAMP," +
			"passwordchangetoken VARCHAR(256)," +
			"passwordchangetokenexpiration TIMESTAMP," +
			"pinnumber INTEGER," +
			"registerip VARCHAR(64)," +
			"lastip VARCHAR(64)," +
			"user_enabled INTEGER NOT NULL DEFAULT '0'," +
			"CONSTRAINT UC_credentials UNIQUE (id, email, username)" +
			")")
	if err := db.Exec(createUsersTable); err.Error != nil {
		log.Panicf("Error creating users table: %v", err)
	}

	// Create snippets table
	createSnippetsTable := fmt.Sprintf(
		"CREATE TABLE IF NOT EXISTS snippets (" +
			"id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT," +
			"owner INTEGER," +
			"ownername VARCHAR(64) NOT NULL," +
			"name VARCHAR(64) NOT NULL," +
			"language VARCHAR(64) NOT NULL," +
			"code VARCHAR(5000) NOT NULL," +
			"tags VARCHAR(256)," +
			"ispublic INTEGER NOT NULL DEFAULT '0'," +
			"CONSTRAINT UC_credentials UNIQUE (id, name)" +
			")")
	if err := db.Exec(createSnippetsTable); err.Error != nil {
		log.Panicf("Error creating snippets table: %v", err)
	}

	// Create snippets table
	createJobsTable := fmt.Sprintf(
		"CREATE TABLE IF NOT EXISTS scheduled_jobs (" +
			"id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT," +
			"cron_exp VARCHAR(64) NOT NULL," +
			"cron_desc VARCHAR(256)," +
			"bot_name VARCHAR(64) NOT NULL," +
			"bot_version INT UNSIGNED NOT NULL," +
			"target_agent VARCHAR(64) NOT NULL," +
			"params VARCHAR(512)," +
			"uuid VARCHAR(128)," +
			"UNIQUE (id)" +
			")")
	if err := db.Exec(createJobsTable); err.Error != nil {
		log.Panicf("Error creating scheduled_jobs table: %v", err)
	}

	// var dbStats []struct {
	// 	Id      uint
	// 	User    string
	// 	Host    string
	// 	Db      string
	// 	Command string
	// 	Time    string
	// 	State   string
	// 	Info    string
	// }
	// db.Raw("SELECT id, user, host, db, state FROM information_schema.processlist;").Scan(&dbStats)
	// log.Printf("Apps connected: %v\n", dbStats)

}
