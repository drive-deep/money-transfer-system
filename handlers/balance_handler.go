package handlers

import (
	"encoding/json"
	"net/http"

	"money-transfer-system/services"
)

// BalanceHandler retrieves the balance of a user
func BalanceHandler(service *services.AccountService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.URL.Query().Get("user")
		if user == "" {
			http.Error(w, "User parameter is required", http.StatusBadRequest)
			return
		}

		balance, err := service.GetBalance(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(map[string]interface{}{
			"user":    user,
			"balance": balance,
		})
	}
}
