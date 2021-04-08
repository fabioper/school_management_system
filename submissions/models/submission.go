package models

import "gorm.io/gorm"

type Submission struct {
	gorm.Model
	Content     string
	StudentID   uint
	ClassroomID uint
	ActivityID  uint
	GradeID     uint
}
