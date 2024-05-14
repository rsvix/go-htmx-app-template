package db

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectMysql() *gorm.DB {
	host := os.Getenv("POSTGRES_HOST")
	// dbName, _ := os.LookupEnv("POSTGRES_DB")
	// user, _ := os.LookupEnv("POSTGRES_USER")
	// password, _ := os.LookupEnv("POSTGRES_PASSWORD")
	appNameDb := os.Getenv("APP_NAME_DB")
	dbConnUrl := os.Getenv("DB_URL")

	driver := strings.Split(dbConnUrl, "://")[0]
	user := strings.Split(strings.Split(dbConnUrl, "://")[1], ":")[0]
	password := strings.SplitN(strings.SplitN(dbConnUrl, ":", 1)[1], "@", 0)[0]
	dbName := strings.SplitN(strings.SplitN(dbConnUrl, "@", 0)[1], ":", 0)[0]
	portStr := strings.SplitN(strings.SplitN(dbConnUrl, ":", 2)[1], "/", 0)[0]
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Panic(err)
	}

	log.Println(driver)
	log.Println(user)
	log.Println(password)
	log.Println(dbName)
	log.Println(portStr)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d application_name='%s' sslmode=disable TimeZone=America/Sao_Paulo", host, user, password, dbName, port, appNameDb)
	database, e := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if e != nil {
		// Create database if it doesn't exists
		if strings.Contains(e.Error(), "does not exist") {
			log.Printf("Creating Mysql database '%s'", dbName)
			dsn = fmt.Sprintf("host=%s user=%s password=%s port=%d application_name='%s' sslmode=disable TimeZone=America/Sao_Paulo", host, user, password, port, appNameDb)
			database, e = gorm.Open(mysql.Open(dsn), &gorm.Config{})
			if e != nil {
				log.Panic(e)
			}
			createDatabaseCommand := fmt.Sprintf("CREATE DATABASE %s", dbName)
			database.Exec(createDatabaseCommand)
		} else {
			log.Panic(e)
		}
	}

	// https://gorm.io/docs/connecting_to_the_database.html#Connection-Pool
	dbConfig, err := database.DB()
	if err != nil {
		log.Panic(err)
	}
	dbConfig.SetMaxIdleConns(5)
	dbConfig.SetMaxOpenConns(15)
	dbConfig.SetConnMaxLifetime(time.Hour)

	// Create users table
	createUsersTable := fmt.Sprintf(
		"CREATE TABLE IF NOT EXISTS users (" +
			"id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT," +
			"email VARCHAR(64) NOT NULL," +
			"username VARCHAR(64) NOT NULL," +
			"firstname VARCHAR(64) NOT NULL," +
			"lastname VARCHAR(64) NOT NULL," +
			"password VARCHAR(128)," +
			"activationtoken VARCHAR(256)," +
			"activationtokenexpiration TIMESTAMP WITH TIME ZONE," +
			"passwordchangetoken VARCHAR(256)," +
			"passwordchangetokenexpiration TIMESTAMP WITH TIME ZONE," +
			"pinnumber INTEGER," +
			"registerip VARCHAR(64)," +
			"lastip VARCHAR(64)," +
			"user_enabled INTEGER NOT NULL DEFAULT '0'," +
			"CONSTRAINT UC_credentials UNIQUE (id, email, username)" +
			")")
	tx2 := database.Exec(createUsersTable)
	if tx2.Error != nil {
		log.Println(tx2.Error)
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
	tx3 := database.Exec(createSnippetsTable)
	if tx3.Error != nil {
		log.Println(tx3.Error)
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
	tx4 := database.Exec(createJobsTable)
	if tx4.Error != nil {
		log.Println(tx4.Error)
	}

	return database
}
