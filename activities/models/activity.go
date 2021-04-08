package models

import (
	"time"

	"gorm.io/gorm"
)

type Activity struct {
	gorm.Model
	Content   string
	TeacherID uint
	Deadline time.Time
}
