package routes

import (
	"net/http"

	"github.com/MiguelLTadeo/Expense-Tracker.git/internal/controllers"
	"gorm.io/gorm"
)

func UserRoutes(db gorm.DB) {
	http.HandleFunc("/user/create", controllers.CreateUser(db))
}
