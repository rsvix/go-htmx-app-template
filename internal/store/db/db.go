package db

import (
	"fmt"
	"log"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(dbName string) *gorm.DB {
	host := "localhost"
	user := "admin"
	password := "123"
	port := 5432
	tableName := "users"

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=America/Sao_Paulo", host, user, password, dbName, port)
	database, e := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if e != nil {
		// Create database if it doesnt exists
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

	createTableCommand := fmt.Sprintf(
		"CREATE TABLE IF NOT EXISTS %s ("+
			"id SERIAL PRIMARY KEY,"+
			"email VARCHAR(64) UNIQUE NOT NULL,"+
			"firstname VARCHAR(64) NOT NULL,"+
			"lastname VARCHAR(64) NOT NULL,"+
			"password VARCHAR(128) NOT NULL,"+
			"activationtoken VARCHAR(128) NOT NULL,"+
			"activationtokenexpiration TIMESTAMP NOT NULL,"+
			"passwordchangetoken VARCHAR(256),"+
			"passwordchangetokenexpiration TIMESTAMP NOT NULL,"+
			"pinnumber INTEGER,"+
			"registerip VARCHAR(64),"+
			"enabled INTEGER NOT NULL DEFAULT '0'"+
			")",
		tableName)
	tx2 := database.Exec(createTableCommand)
	if tx2.Error != nil {
		log.Println(tx2.Error)
	}

	return database
}
