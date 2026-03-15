package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string    `gorm:"unique;not null" json:"email"`
	Password string    `json:"password"`
	Expenses []Expense `gorm:"constraint:OnDelete:CASCADE;" json:"expenses"`
}
