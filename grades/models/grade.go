package models

import "gorm.io/gorm"

type Grade struct {
	gorm.Model
	SubmissionID uint
	TeacherID    uint
	Grade        float32
}
