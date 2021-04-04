package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("submissions.db"), &gorm.Config{})

	if err != nil {
		panic("[Submissions Service] Failed to connect to database.")
	}

	database.AutoMigrate(&Submission{})

	DB = database
}
