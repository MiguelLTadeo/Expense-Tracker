package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string    `gorm:"unique" json:"email"`
	Password string    `json:"password"`
	Expenses []Expense `json:"expenses"`
}

func GetAll(db *gorm.DB) ([]User, error) {
	var users []User
	err := db.Model(&User{}).Preload("Expenses").Find(&users).Error
	return users, err
}
