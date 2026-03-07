package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MiguelLTadeo/Expense-Tracker.git/internal/models"
	"github.com/MiguelLTadeo/Expense-Tracker.git/internal/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateUserHandler(db gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {

			var userJson struct {
				Email            string `gorm:"unique" json:"email"`
				Password         string `json:"password"`
				Confirm_password string `json:"confirm_password"`
			}

			err := json.NewDecoder(r.Body).Decode(&userJson)

			if err != nil {
				http.Error(w, "Error reading Json: "+err.Error(), http.StatusBadRequest)

				return
			}

			if userJson.Password == "" || userJson.Email == "" || userJson.Confirm_password == "" {
				http.Error(w, "Missing fields", http.StatusBadRequest)

				return
			}

			if userJson.Password != userJson.Confirm_password {
				http.Error(w, "Divergent passwords", http.StatusUnprocessableEntity)

				return
			}

			bytes, err := bcrypt.GenerateFromPassword([]byte(userJson.Password), 10)

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			userJson.Password = string(bytes)

			newUser := models.User{}

			newUser.Email = userJson.Email

			newUser.Password = userJson.Password

			result := db.Create(&newUser)

			if result.Error != nil {
				http.Error(w, result.Error.Error(), http.StatusInternalServerError)

				return
			}

			w.Header().Set("Content-Type", "application/json")

			w.WriteHeader(http.StatusAccepted)

			fmt.Fprintf(w, "UsersStatusAccepted, updated rows %d", result.RowsAffected)

		} else {

			w.WriteHeader(http.StatusMethodNotAllowed)

		}
	}
}

func LoginUserhandler(db gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			loginUser := models.User{}

			err := json.NewDecoder(r.Body).Decode(&loginUser)

			if loginUser.Password == "" || loginUser.Email == "" {
				http.Error(w, "Missing fields", http.StatusBadRequest)

				return
			}

			if err != nil {
				http.Error(w, "Error reading Json: "+err.Error(), http.StatusBadRequest)

				return
			}

			userExists := models.User{}

			result := db.Where("email = ?", loginUser.Email).First(&userExists)

			if result.Error != nil {
				http.Error(w, "Password or Email incorrect", http.StatusUnauthorized)

				return
			}

			err = bcrypt.CompareHashAndPassword([]byte(userExists.Password), []byte(loginUser.Password))

			if err != nil {
				http.Error(w, "Password or Email incorrect", http.StatusUnauthorized)

				return
			} else {

				token, err := utils.CreateToken(loginUser.Email)

				if err != nil {
					http.Error(w, "Internal error", http.StatusInternalServerError)

					return
				}

				newTokenResponse := utils.TokenResponse{
					Token: token,
				}

				err = json.NewEncoder(w).Encode(newTokenResponse)

				if err != nil {
					http.Error(w, "Error in response", http.StatusInternalServerError)

					return
				}

				w.Header().Set("Content-Type", "application/json")

				w.WriteHeader(http.StatusOK)
			}

		} else {

			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func DeleteUserHandler(db gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "DELETE" {
			//deleteUser := models.User{}

			//id, err := strconv.Atoi(r.FormValue("id"))

			//if err != nil {
			//http.Error(w, "Invalid id: "+err.Error(), http.StatusBadRequest)

			//return
			//}

			//deleteUser.ID
		} else {

			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}
