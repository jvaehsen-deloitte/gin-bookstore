package models

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("sqlite3", "./data/test.db")

	if err != nil {
		log.Fatal(err.Error())
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Book{})

	DB = database
}
