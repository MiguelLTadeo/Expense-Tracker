package main

import (
	"fmt"

	databaseConn "github.com/MiguelLTadeo/Expense-Tracker.git/internal/database-conn"
)

func main() {
	db := databaseConn.Init()
	fmt.Println(db)
}
