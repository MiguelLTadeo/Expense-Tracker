package routes

import (
	"fmt"
	"io"
	"net/http"

	"github.com/MiguelLTadeo/Expense-Tracker.git/internal/service/controllers"
	"github.com/didip/tollbooth"
	"gorm.io/gorm"
)

func UserRoutes(db *gorm.DB) {
	limiter := tollbooth.NewLimiter(5, nil) // 5 req/s
	limiter.SetIPLookups([]string{"X-Forwarded-For", "RemoteAddr"})

	http.Handle("/hello", tollbooth.LimitFuncHandler(limiter, getHello))

	http.Handle("/user/create", tollbooth.LimitFuncHandler(limiter, controllers.CreateUserHandler(db)))
	http.Handle("/user/delete", tollbooth.LimitFuncHandler(limiter, controllers.DeleteUserHandler(db)))
	http.Handle("/user/login", tollbooth.LimitFuncHandler(limiter, controllers.LoginUserhandler(db)))
	http.Handle("/user/update", tollbooth.LimitFuncHandler(limiter, controllers.UpdateUserHandler(db)))
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}
