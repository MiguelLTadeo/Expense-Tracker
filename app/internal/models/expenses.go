package models

import "gorm.io/gorm"

type Expense struct {
	gorm.Model
	Description string `json:"description"`
	Value       string `json:"value"`
	UserID      uint   `json:"user_id"`
}
