package routes

import (
	"net/http"

	"github.com/MiguelLTadeo/Expense-Tracker.git/internal/service/controllers"
	"github.com/didip/tollbooth"
	"gorm.io/gorm"
)

func ExpenseRoutes(db *gorm.DB) {
	limiter := tollbooth.NewLimiter(5, nil) // 5 req/s
	limiter.SetIPLookups([]string{"X-Forwarded-For", "RemoteAddr"})

	http.Handle("/expense/create", tollbooth.LimitFuncHandler(limiter, controllers.CreateExpenseHandler(db)))
	http.Handle("/expense/delete", tollbooth.LimitFuncHandler(limiter, controllers.DeleteExpenseHandler(db)))
	http.Handle("/expense/get", tollbooth.LimitFuncHandler(limiter, controllers.GetExpenseHandler(db)))
	http.Handle("/expense/list", tollbooth.LimitFuncHandler(limiter, controllers.ListExpenseHandler(db)))
	http.Handle("/expense/update", tollbooth.LimitFuncHandler(limiter, controllers.UpdateExpenseHandler(db)))
}
