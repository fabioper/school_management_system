package models

import "gorm.io/gorm"

type Submission struct {
	gorm.Model
	Content     string
	ActivityId  uint
	ClassroomId uint
	StudentId   uint
	Grade       float32
}
