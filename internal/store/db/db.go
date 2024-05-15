package db

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dbConnUrl := os.Getenv("DB_URL")
	driver := strings.Split(dbConnUrl, "://")[0]

	if driver == "mysql" {
		dsn1 := strings.Split(dbConnUrl, "://")[1]
		dsn2 := strings.Split(dsn1, "/")[0]
		dbName := strings.Split(dsn1, "/")[1]

		database, e := gorm.Open(mysql.Open(dsn1+"?charset=utf8&parseTime=true"), &gorm.Config{})
		if e != nil {
			log.Println(e)
			if strings.Contains(e.Error(), "does not exist") {
				log.Printf("Creating Mysql database '%s'", dbName)
				database, e = gorm.Open(mysql.Open(dsn2+"?charset=utf8&parseTime=true"), &gorm.Config{})
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

		MigrateDb(database)
		return database
	} else if driver == "postgres" {
		return nil
	}
	return nil
}
