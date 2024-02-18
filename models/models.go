package models

import "gorm.io/gorm"

type agency2 struct {
	gorm.Model
	agency string
	car    string
}
