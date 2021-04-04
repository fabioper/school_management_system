package models

import "gorm.io/gorm"

type Activity struct {
	gorm.Model
	Content string
}
