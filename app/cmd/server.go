package main

import (
	"fmt"
	"net/http"

	databaseConn "github.com/MiguelLTadeo/Expense-Tracker.git/internal/database-conn"
	"github.com/MiguelLTadeo/Expense-Tracker.git/internal/routes"
)

func main() {
	db := databaseConn.Init()
	routes.UserRoutes(db)
	routes.ExpenseRoutes(db)
	http.ListenAndServe(":8080", nil)
	fmt.Println(db)
}
