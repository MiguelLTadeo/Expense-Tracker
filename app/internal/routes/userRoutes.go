package routes

import (
	"net/http"

	"github.com/MiguelLTadeo/Expense-Tracker.git/internal/controllers"
	"gorm.io/gorm"
)

func UserRoutes(db gorm.DB) {
	http.HandleFunc("/user/create", controllers.CreateUser(db))
}

func RouteHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is my API!"))
}
