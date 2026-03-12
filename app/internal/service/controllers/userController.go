package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MiguelLTadeo/Expense-Tracker.git/internal/middleware"
	"github.com/MiguelLTadeo/Expense-Tracker.git/internal/service/models"
	"github.com/MiguelLTadeo/Expense-Tracker.git/internal/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

//struct para pegar o Json do Body

type UserJson struct {
	Email            string `gorm:"unique;not null" json:"email"`
	Password         string `json:"password"`
	Confirm_password string `json:"confirm_password"`
}

func CreateUserHandler(db gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {

			w.Header().Set("Content-Type", "application/json")

			userJson := UserJson{}

			err := json.NewDecoder(r.Body).Decode(&userJson)

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)

				err = json.NewEncoder(w).Encode(utils.ErrorResponse{
					Error: "Internal error",
				})

				return
			}

			if userJson.Password == "" || userJson.Email == "" || userJson.Confirm_password == "" {
				w.WriteHeader(http.StatusBadRequest)

				err = json.NewEncoder(w).Encode(utils.ErrorResponse{
					Error: "Missing fields",
				})

				return
			}

			if userJson.Password != userJson.Confirm_password {
				w.WriteHeader(http.StatusUnprocessableEntity)

				err = json.NewEncoder(w).Encode(utils.ErrorResponse{
					Error: "Divergent passwords",
				})

				return
			}

			bytes, err := bcrypt.GenerateFromPassword([]byte(userJson.Password), 10)

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)

				err = json.NewEncoder(w).Encode(utils.ErrorResponse{
					Error: "Internal error",
				})

				return
			}

			userJson.Password = string(bytes)

			newUser := models.User{}

			newUser.Email = userJson.Email

			newUser.Password = userJson.Password

			if db.Where("email = ?", newUser.Email).First(&models.User{}).Error == nil {
				w.WriteHeader(http.StatusConflict)

				err = json.NewEncoder(w).Encode(utils.ErrorResponse{
					Error: "Email in use",
				})
				return
			}

			result := db.Create(&newUser)

			if result.Error != nil {
				w.WriteHeader(http.StatusInternalServerError)

				err = json.NewEncoder(w).Encode(utils.ErrorResponse{
					Error: "Internal error",
				})
				return
			}

			w.WriteHeader(http.StatusCreated)

			err = json.NewEncoder(w).Encode(utils.SucessResponse{
				Message: "User created",
			})

		} else {

			w.WriteHeader(http.StatusMethodNotAllowed)

		}
	}
}

func LoginUserhandler(db gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {

			w.Header().Set("Content-Type", "application/json")

			loginUser := models.User{}

			err := json.NewDecoder(r.Body).Decode(&loginUser)

			if loginUser.Password == "" || loginUser.Email == "" {
				w.WriteHeader(http.StatusBadRequest)

				err = json.NewEncoder(w).Encode(utils.ErrorResponse{
					Error: "Missing fields",
				})

				return
			}

			if err != nil {
				w.WriteHeader(http.StatusBadRequest)

				err = json.NewEncoder(w).Encode(utils.ErrorResponse{
					Error: "Error reading Json",
				})

				return
			}

			userExists := models.User{}

			result := db.Where("email = ?", loginUser.Email).First(&userExists)

			if result.Error != nil {
				w.WriteHeader(http.StatusUnauthorized)

				err = json.NewEncoder(w).Encode(utils.ErrorResponse{
					Error: "Password or Email incorrect",
				})

				return
			}

			err = bcrypt.CompareHashAndPassword([]byte(userExists.Password), []byte(loginUser.Password))

			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)

				err = json.NewEncoder(w).Encode(utils.ErrorResponse{
					Error: "Password or Email incorrect",
				})

				return
			} else {

				token, err := utils.CreateToken(loginUser.Email)

				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)

					err = json.NewEncoder(w).Encode(utils.ErrorResponse{
						Error: "Internal error",
					})

					return
				}

				w.WriteHeader(http.StatusAccepted)

				err = json.NewEncoder(w).Encode(utils.TokenResponse{
					Token: token,
				})

				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)

					err = json.NewEncoder(w).Encode(utils.ErrorResponse{
						Error: "Internal error",
					})

					return
				}

			}

		} else {

			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func DeleteUserHandler(db gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "DELETE" {

			w.Header().Set("Content-Type", "application/json")

			email, err := middleware.AuthToken(w, r, db)

			fmt.Println(email)

			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)

				err = json.NewEncoder(w).Encode(utils.ErrorResponse{
					Error: "Invalid token",
				})

				return
			} else {

				db.Unscoped().Delete(&models.User{}, "email = ?", email)

				w.WriteHeader(http.StatusOK)

				err = json.NewEncoder(w).Encode(utils.SucessResponse{
					Message: "User deleted",
				})
				return

			}

		} else {

			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func UpdateUserHandler(db gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "PUT" {

			w.Header().Set("Content-Type", "application/json")

			updateUser := UserJson{}

			err := json.NewDecoder(r.Body).Decode(&updateUser)

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)

				err = json.NewEncoder(w).Encode(utils.ErrorResponse{
					Error: "Internal error",
				})

				return
			}

			email, err := middleware.AuthToken(w, r, db)

			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)

				err = json.NewEncoder(w).Encode(utils.ErrorResponse{
					Error: "Invalid token",
				})

				return
			} else {
				userExists := models.User{}

				db.Where(&models.User{Email: email}).First(&userExists)

				if userExists.Email == "" {
					w.WriteHeader(http.StatusNotFound)

					err = json.NewEncoder(w).Encode(utils.ErrorResponse{
						Error: "User not found",
					})
				} else {
					if db.Where("email = ?", updateUser.Email).First(&models.User{}).Error == nil {
						w.WriteHeader(http.StatusConflict)

						err = json.NewEncoder(w).Encode(utils.ErrorResponse{
							Error: "Email in use",
						})
						return
					} else {
						if updateUser.Email != "" {
							userExists.Email = updateUser.Email
						}

						db.Save(&userExists)

						w.WriteHeader(http.StatusOK)

						err = json.NewEncoder(w).Encode(utils.SucessResponse{
							Message: "User updated",
						})
						return
					}
				}

			}
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)

		}
	}
}
