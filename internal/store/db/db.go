package db

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	host, _ := os.LookupEnv("POSTGRES_HOST")
	dbName, _ := os.LookupEnv("POSTGRES_DB")
	user, _ := os.LookupEnv("POSTGRES_USER")
	password, _ := os.LookupEnv("POSTGRES_PASSWORD")

	p, _ := os.LookupEnv("POSTGRES_PORT")
	port, err := strconv.Atoi(p)
	if err != nil {
		log.Panic(err)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=America/Sao_Paulo", host, user, password, dbName, port)
	database, e := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if e != nil {
		// Create database if it doesn't exists
		if strings.Contains(e.Error(), "does not exist") {
			log.Printf("Creating Postgres database '%s'", dbName)
			dsn = fmt.Sprintf("host=%s user=%s password=%s port=%d sslmode=disable TimeZone=America/Sao_Paulo", host, user, password, port)
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

	tableName := "users"

	createTableCommand := fmt.Sprintf(
		"CREATE TABLE IF NOT EXISTS %s ("+
			"id SERIAL PRIMARY KEY,"+
			"email VARCHAR(64) UNIQUE NOT NULL,"+
			"firstname VARCHAR(64) NOT NULL,"+
			"lastname VARCHAR(64) NOT NULL,"+
			"password VARCHAR(128) NOT NULL,"+
			"activationtoken VARCHAR(256),"+
			"activationtokenexpiration TIMESTAMP,"+
			"passwordchangetoken VARCHAR(256),"+
			"passwordchangetokenexpiration TIMESTAMP,"+
			"pinnumber INTEGER,"+
			"registerip VARCHAR(64),"+
			"lastip VARCHAR(64),"+
			"enabled INTEGER NOT NULL DEFAULT '0'"+
			")",
		tableName)
	tx2 := database.Exec(createTableCommand)
	if tx2.Error != nil {
		log.Println(tx2.Error)
	}

	return database
}
