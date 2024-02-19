package models

import "gorm.io/gorm"

type Fact struct {
	gorm.Model
	Question string `json:"question" gorm:"text;not null;default:null"`
	Answer   string `json:"answer" gorm:"text;not null;default:null"`
}

type Agency struct {
	gorm.Model
	Name   string `json:"name" gorm:"text;not null;default:null"`
	Adress string `json:"adress" gorm:"text;not null;default:null"`
	Phone  string `json:"phone" gorm:"text;not null;default:null"`
	Email  string `json:"email" gorm:"text;not null;default:null"`
}

type Car struct {
	gorm.Model
	Registration_number string `json:"registration_number" gorm:"text;not null;default:null"`
	Brand               string `json:"brand" gorm:"text;not null;default:null"`
	Color               string `json:"color" gorm:"text;not null;default:null"`
	Year                string `json:"year" gorm:"text;not null;default:null"`
}
