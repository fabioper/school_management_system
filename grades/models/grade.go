package models

import "gorm.io/gorm"

type Grade struct {
	gorm.Model
	StudentID uint
	ActivityID uint
	ClassroomID uint
	Grade uint
}
