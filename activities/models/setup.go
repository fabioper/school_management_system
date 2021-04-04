package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("activities.db"), &gorm.Config{})

	if err != nil {
		panic("[Activities Service] Failed to connect to database.")
	}

	database.AutoMigrate(&Activity{})

	DB = database
}
