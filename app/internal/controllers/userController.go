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

			newUser := models.User{}

			err := json.NewDecoder(r.Body).Decode(&newUser)

			if err != nil {
				http.Error(w, "Error reading Json: "+err.Error(), http.StatusBadRequest)
			}

			newUser.Password = utils.GetMD5Hash(newUser.Password)

			result := db.Create(&newUser)

			if result.Error != nil {
				http.Error(w, "Error creating User: "+result.Error.Error(), http.StatusInternalServerError)
			}

			w.Header().Set("Content-Type", "application/json")

			w.WriteHeader(http.StatusCreated)

			fmt.Fprintf(w, "User created, updated rows %d", result.RowsAffected)
		}
	}
}

func DeleteUser(db gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
