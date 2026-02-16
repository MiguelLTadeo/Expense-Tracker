package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string
	Expenses []Expense
}

func GetAll(db *gorm.DB) ([]User, error) {
	var users []User
	err := db.Model(&User{}).Preload("Expenses").Find(&users).Error
	return users, err
}
