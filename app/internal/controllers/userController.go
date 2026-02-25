package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MiguelLTadeo/Expense-Tracker.git/internal/models"
	"github.com/MiguelLTadeo/Expense-Tracker.git/internal/utils"
	"gorm.io/gorm"
)

func CreateUser(db gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			fmt.Println(r.FormValue("email"))
			email := r.FormValue("email")
			password := r.FormValue("password")
			password = utils.GetMD5Hash(password)
			fmt.Println(email)

			user := models.User{}
			err := json.NewDecoder(r.Body).Decode(&user)

			fmt.Println(r.FormValue("password"))

			fmt.Println(user)
			//result := db.Create(&user)

			//print(result)
		}
	}
}

func DeleteUser(db gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
