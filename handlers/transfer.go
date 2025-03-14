package handlers

import (
	"encoding/json"
	"money-transfer-system/models"
	"money-transfer-system/services"
	"net/http"
)

// TransferHandler handles money transfer HTTP requests
func TransferHandler(service *services.AccountService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req models.TransferRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		if err := service.Transfer(req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Transfer successful"})
	}
}
