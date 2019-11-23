package db

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/victor-nach/nigerian-senators/api/models"
)

var db *gorm.DB

func init() {
	dbURL, ok := os.LookupEnv("DATABASE_URL")
	if !ok {
		dbURL = "postgres://xrqfqmoo:5jTTWJqZsQoWXCsUl5MVwZ2CO2FXiycB@salt.db.elephantsql.com:5432/xrqfqmoo"
	}
	var err error
	db, err = gorm.Open("postgres", dbURL)

	if err != nil {
		log.Println("Could not connect to the db")
		log.Fatal(err)
	}
	log.Println("Succesfully connected to the db")
	defer db.Close()

	db.Debug().AutoMigrate(&models.Senators{})
}

func Db() *gorm.DB {
	return db
}