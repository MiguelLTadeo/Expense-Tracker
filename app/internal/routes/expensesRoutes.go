package routes

import (
	"net/http"

	"github.com/MiguelLTadeo/Expense-Tracker.git/internal/service/controllers"
	"gorm.io/gorm"
)

func ExpenseRoutes(db *gorm.DB) {
	http.HandleFunc("/expense/create", controllers.CreateExpenseHandler(db))
	http.HandleFunc("/expense/delete", controllers.DeleteExpenseHandler(db))
	http.HandleFunc("/expense/get", controllers.GetExpenseHandler(db))
	http.HandleFunc("/expense/list", controllers.ListExpenseHandler(db))

}
