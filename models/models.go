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
	Name      string `json:"name" gorm:"text;not null;default:null"`
	Address   string `json:"adress" gorm:"text;default:null"`
	Phone     string `json:"phone" gorm:"text;not null;default:null"`
	Email     string `json:"email" gorm:"text;default:null"`
	Cars      []Car
	Clients   []Client
	Employees []Employee
}

type Car struct {
	gorm.Model
	Registration_number string `json:"registration_number" gorm:"text;not null;default:null"`
	Brand               string `json:"brand" gorm:"text;not null;default:null"`
	Color               string `json:"color" gorm:"text;not null;default:null"`
	Year                string `json:"year" gorm:"text;not null;default:null"`
	AgencyID            uint   `json:"agency_id" gorm:"default:null"`
}

type Person struct {
	gorm.Model
	Name    string `json:"name" gorm:"text;not null;default:null"`
	Email   string `json:"email" gorm:"text;not null;default:null"`
	Address string `json:"address" gorm:"text;default:null"`
	Phone   string `json:"phone" gorm:"text;not null;default:null"`
}

type Client struct {
	gorm.Model
	Person          // Embedding Person struct to achieve inheritance-like behavior
	AgencyID uint   `json:"agency_id" gorm:"default:null"`
	Agency   Agency `gorm:"foreignkey:AgencyID"`
}

type Employee struct {
	gorm.Model
	Person          // Embedding Person struct to achieve inheritance-like behavior
	Salary   string `json:"salary" gorm:"text;null;default:null"`
	AgencyID uint   `json:"agency_id" gorm:"default:null"`
	Agency   Agency `gorm:"foreignkey:AgencyID"`
}

type Reservation struct {
	gorm.Model
	Date       time.Time `json:"date" gorm:"type:date;not null"`
	Duration   int       `json:"duration" gorm:"not null"`
	StartDate  time.Time `json:"start_date" gorm:"type:date;not null"`
	EndDate    time.Time `json:"end_date" gorm:"type:date;not null"`
	ClientID   uint      `json:"client_id" gorm:"default:null"`
	CarID      uint      `json:"car_id" gorm:"default:null"`
	Client     Client    `gorm:"foreignkey:ClientID"`
	Car        Car       `gorm:"foreignkey:CarID"`
	ContractID uint      `json:"contract_id" gorm:"default:null"`
	Contract   Contract  `gorm:"foreignkey:ContractID"`
}

type Contract struct {
	gorm.Model
	Date         time.Time `json:"date" gorm:"type:date;not null"`
	Duration     int       `json:"duration" gorm:"not null"`
	StartDate    time.Time `json:"start_date" gorm:"type:date;not null"`
	EndDate      time.Time `json:"end_date" gorm:"type:date;not null"`
	TotalPrice   float64   `json:"totalprice" gorm:"not null"`
	ClientID     uint      `json:"client_id" gorm:"default:null"`
	Client       Client    `gorm:"foreignkey:ClientID"`
	Reservations []Reservation
}

type Bill struct {
	gorm.Model
	Date       time.Time `json:"date" gorm:"type:date;not null"`
	Duration   int       `json:"duration" gorm:"not null"`
	TotalPrice float64   `json:"totalprice" gorm:"not null"`
	ClientID   uint      `json:"client_id" gorm:"default:null"`
	Client     Client    `gorm:"foreignkey:ClientID"`
	Contracts  []Contract
}

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"text;not null;default:null"`
	Email    string `json:"email" gorm:"text;not null;default:null"`
	Password string `json:"password" gorm:"text;not null;default:null"`
}
