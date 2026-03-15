package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/MiguelLTadeo/Expense-Tracker.git/internal/middleware"
	"github.com/MiguelLTadeo/Expense-Tracker.git/internal/service/models"
	"github.com/MiguelLTadeo/Expense-Tracker.git/internal/utils"
	"gorm.io/gorm"
)

func CreateExpenseHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {

			w.Header().Set("Content-Type", "application/json")

			email, err := middleware.AuthToken(w, r, db)

			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)

				err = json.NewEncoder(w).Encode(utils.ErrorResponse{
					Error: "Invalid token",
				})

				return
			}

			expenseJson := models.Expense{}

			err = json.NewDecoder(r.Body).Decode(&expenseJson)

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)

				err = json.NewEncoder(w).Encode(utils.ErrorResponse{
					Error: "Internal error",
				})

				return
			}

			if expenseJson.Description == "" || expenseJson.Value == 0 {
				w.WriteHeader(http.StatusBadRequest)

				err = json.NewEncoder(w).Encode(utils.ErrorResponse{
					Error: "Missing fields",
				})

				return
			}

			expensesUser := models.User{}

			db.Where(&models.User{Email: email}).First(&expensesUser)

			if expensesUser.Email == "" {
				w.WriteHeader(http.StatusNotFound)

				err = json.NewEncoder(w).Encode(utils.ErrorResponse{
					Error: "User not found",
				})
				return
			}

			expenseJson.UserID = expensesUser.ID

			result := db.Create(&expenseJson)

			if result.Error != nil {
				w.WriteHeader(http.StatusInternalServerError)

				err = json.NewEncoder(w).Encode(utils.ErrorResponse{
					Error: "Internal error",
				})
				return
			}

			w.WriteHeader(http.StatusCreated)

			err = json.NewEncoder(w).Encode(utils.SucessResponse{
				Message: "Expense created",
			})

		} else {

			w.WriteHeader(http.StatusMethodNotAllowed)

		}
	}
}

func DeleteExpenseHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "DELETE" {

			w.Header().Set("Content-Type", "application/json")

			email, err := middleware.AuthToken(w, r, db)

			expenseId := r.FormValue("id")

			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)

				err = json.NewEncoder(w).Encode(utils.ErrorResponse{
					Error: "Invalid token",
				})

				return
			}

			db.Unscoped().Delete(&models.Expense{}, "id = ? AND user_id = ?",
				expenseId,
				db.Model(&models.User{}).Select("id").Where("email = ?", email),
			)

			w.WriteHeader(http.StatusOK)

			err = json.NewEncoder(w).Encode(utils.SucessResponse{
				Message: "Expense deleted",
			})

			return

		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func UpdateExpenseHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "PUT" {

			w.Header().Set("Content-Type", "application/json")

			updateExpense := models.Expense{}

			err := json.NewDecoder(r.Body).Decode(&updateExpense)

			expenseId := r.FormValue("id")

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
			}
			userExists := models.User{}

			db.Where(&models.User{Email: email}).First(&userExists)

			if userExists.Email == "" {
				w.WriteHeader(http.StatusNotFound)

				err = json.NewEncoder(w).Encode(utils.ErrorResponse{
					Error: "User not found",
				})
				return
			}

			expenseExist := models.Expense{}

			if db.Where("id = ? AND user_id = ?", expenseId, userExists.ID).First(&expenseExist).Error != nil {
				w.WriteHeader(http.StatusNotFound)

				err = json.NewEncoder(w).Encode(utils.ErrorResponse{
					Error: "Task dont exist",
				})
				return
			}

			if updateExpense.Value != 0 {
				expenseExist.Value = updateExpense.Value
			}
			if updateExpense.Description != "" {
				expenseExist.Description = updateExpense.Description

			}

			if err := db.Save(&expenseExist).Error; err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(utils.ErrorResponse{Error: "Error updating expense"})
				return
			}

			w.WriteHeader(http.StatusOK)

			err = json.NewEncoder(w).Encode(utils.SucessResponse{
				Message: "User updated",
			})
			return

		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)

		}
	}
}

func GetExpenseHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {

			w.Header().Set("Content-Type", "application/json")

			email, err := middleware.AuthToken(w, r, db)

			expenseId := r.FormValue("id")

			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)

				err = json.NewEncoder(w).Encode(utils.ErrorResponse{
					Error: "Invalid token",
				})

				return
			}

			expense := models.Expense{}

			db.Where("id = ? AND user_id = (?)",
				expenseId,
				db.Model(&models.User{}).Select("id").Where("email = ?", email),
			).First(&expense)

			w.WriteHeader(http.StatusOK)

			err = json.NewEncoder(w).Encode(expense)

		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func ListExpenseHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {

			w.Header().Set("Content-Type", "application/json")

			email, err := middleware.AuthToken(w, r, db)

			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)

				err = json.NewEncoder(w).Encode(utils.ErrorResponse{
					Error: "Invalid token",
				})

				return
			}

			search := r.URL.Query().Get("search")
			orderBy := r.URL.Query().Get("order_by")
			if orderBy == "" {
				orderBy = "created_at"
			}

			var total int64

			var expenses []models.Expense

			query := db.Model(&models.Expense{}).Where("user_id = (?)",
				db.Model(&models.User{}).Select("id").Where("email = ?", email),
			)

			if search != "" {
				query = query.Where("description ILIKE ?", "%"+search+"%")
			}

			query.Session(&gorm.Session{}).Count(&total)

			var info utils.PageInfo

			query.Scopes(utils.Paginate(r, &info)).Order(orderBy).Find(&expenses)

			w.WriteHeader(http.StatusOK)

			err = json.NewEncoder(w).Encode(map[string]any{
				"data":      expenses,
				"page":      info.Page,
				"page_size": info.PageSize,
				"total":     total,
			})

		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}
