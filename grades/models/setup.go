package models

import (
	"errors"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() error {
	database, err := gorm.Open(sqlite.Open("grades.db"), &gorm.Config{})

	if err != nil {
		return errors.New("failed to connect to database")
	}

	if err := database.AutoMigrate(&Grade{}); err != nil {
		return errors.New("failed to migrate database")
	}

	DB = database
	return nil
}