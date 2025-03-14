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

		// Decode the request body into TransferRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			respondWithJSON(w, http.StatusBadRequest, "Invalid request")
			return
		}

		// Perform the transfer
		if err := service.Transfer(req); err != nil {
			respondWithJSON(w, http.StatusBadRequest, err.Error())
			return
		}

		// Send success response
		respondWithJSON(w, http.StatusOK, "Transfer successful")
	}
}

func respondWithJSON(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{
		"status":  http.StatusText(status),
		"message": message,
	})
}