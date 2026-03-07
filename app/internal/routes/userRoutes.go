package routes

import (
	"net/http"

	"github.com/MiguelLTadeo/Expense-Tracker.git/internal/controllers"
	"gorm.io/gorm"
)

func UserRoutes(db gorm.DB) {
	http.HandleFunc("/user/create", controllers.CreateUserHandler(db))
	http.HandleFunc("/user/delete", controllers.DeleteUserHandler(db))
	http.HandleFunc("/user/login", controllers.LoginUserhandler(db))

}
