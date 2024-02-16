package models

import "gorm.io/gorm"

type agency struct {
    gorm.Model
    agency string
    car   string
}