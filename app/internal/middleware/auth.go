package middleware

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/MiguelLTadeo/Expense-Tracker.git/internal/service/models"
	"github.com/MiguelLTadeo/Expense-Tracker.git/internal/utils"
	"gorm.io/gorm"
)

func AuthToken(w http.ResponseWriter, r *http.Request, db gorm.DB) (string, error) {
	authHeader := r.Header.Get("Authorization")

	if authHeader == "" {
		w.WriteHeader(http.StatusUnauthorized)

		err := json.NewEncoder(w).Encode(utils.ErrorResponse{
			Error: "Missing token",
		})

		return "", err
	}

	if !strings.HasPrefix(authHeader, "Bearer ") {
		w.WriteHeader(http.StatusUnauthorized)

		err := json.NewEncoder(w).Encode(utils.ErrorResponse{
			Error: "Invalid authentication",
		})

		return "", err
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	email, err := utils.VerifyToken(tokenString)

	if err != nil {
		return "", err
	} else {
		err = db.Where("email = ?", email).First(&models.User{}).Error
		if err != nil {
			return "", err
		}
		return email, nil
	}

}
