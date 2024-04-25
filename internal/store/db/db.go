package db

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	host, _ := os.LookupEnv("POSTGRES_HOST")
	dbName, _ := os.LookupEnv("POSTGRES_DB")
	user, _ := os.LookupEnv("POSTGRES_USER")
	password, _ := os.LookupEnv("POSTGRES_PASSWORD")
	appNameDb, _ := os.LookupEnv("APP_NAME_DB")

	p, _ := os.LookupEnv("POSTGRES_PORT")
	port, err := strconv.Atoi(p)
	if err != nil {
		log.Panic(err)
	}

	usersTableName := "users"
	snippetsTableName := "snippets"

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d application_name='%s' sslmode=disable TimeZone=America/Sao_Paulo", host, user, password, dbName, port, appNameDb)
	database, e := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if e != nil {
		// Create database if it doesn't exists
		if strings.Contains(e.Error(), "does not exist") {
			log.Printf("Creating Postgres database '%s'", dbName)
			dsn = fmt.Sprintf("host=%s user=%s password=%s port=%d application_name='%s' sslmode=disable TimeZone=America/Sao_Paulo", host, user, password, port, appNameDb)
			database, e = gorm.Open(postgres.Open(dsn), &gorm.Config{})
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

	createUsersTable := fmt.Sprintf(
		"CREATE TABLE IF NOT EXISTS %s ("+
			"id SERIAL PRIMARY KEY,"+
			"email VARCHAR(64) UNIQUE NOT NULL,"+
			"username VARCHAR(64) UNIQUE NOT NULL,"+
			"firstname VARCHAR(64) NOT NULL,"+
			"lastname VARCHAR(64) NOT NULL,"+
			"password VARCHAR(128) NOT NULL,"+
			"activationtoken VARCHAR(256),"+
			"activationtokenexpiration TIMESTAMP WITH TIME ZONE,"+
			"passwordchangetoken VARCHAR(256),"+
			"passwordchangetokenexpiration TIMESTAMP WITH TIME ZONE,"+
			"pinnumber INTEGER,"+
			"registerip VARCHAR(64),"+
			"lastip VARCHAR(64),"+
			"enabled INTEGER NOT NULL DEFAULT '0'"+
			")",
		usersTableName)
	tx2 := database.Exec(createUsersTable)
	if tx2.Error != nil {
		log.Println(tx2.Error)
	}

	createSnippetsTable := fmt.Sprintf(
		"CREATE TABLE IF NOT EXISTS %s ("+
			"id SERIAL PRIMARY KEY,"+
			"owner INTEGER,"+
			"ownername VARCHAR(64) NOT NULL,"+
			"name VARCHAR(64) UNIQUE NOT NULL,"+
			"language VARCHAR(64) NOT NULL,"+
			"code VARCHAR(5000) NOT NULL,"+
			"tags VARCHAR(256),"+
			"ispublic INTEGER NOT NULL DEFAULT '0'"+
			")",
		snippetsTableName)
	tx3 := database.Exec(createSnippetsTable)
	if tx3.Error != nil {
		log.Println(tx3.Error)
	}

	return database
}
