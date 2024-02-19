package models

import (
	"time"

	"gorm.io/gorm"
)

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
	AgencyID            uint   `json:"agency_id"`           // Foreign key reference to Agency
	Agency              Agency `gorm:"foreignKey:AgencyID"` // Define the relationship
}

type Person struct {
	gorm.Model
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

type Client struct {
	gorm.Model
	Person          // Embedding Person struct to achieve inheritance-like behavior
	AgencyID uint   // Foreign key reference to Agency
	Agency   Agency `gorm:"foreignkey:AgencyID"` // Define the relationship
}

type Reservation struct {
	gorm.Model
	Date       time.Time `json:"date"`
	Duration   int       `json:"email"`
	Start_date time.Time `json:"start_date"`
	End_date   time.Time `json:"end_date"`
}
