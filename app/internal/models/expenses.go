package models

import "gorm.io/gorm"

type Expense struct {
	gorm.Model
	Description string
	Value       string
	UserID      uint
}
