package models

import "gorm.io/gorm"

type agency struct {
	gorm.Model
	name   string `json:"name" gorm:"text;not null;default:null`
	adress string `json:"adress" gorm:"text;not null;default:null`
}
