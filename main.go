package main

import (
	"log"
	"net/http"
	"money-transfer-system/handlers"
	"money-transfer-system/services"
	"money-transfer-system/storage"
)

func main() {
	// Initialize account service with in-memory storage
	accountService := services.NewAccountService(storage.InitializeDatabase())

	// Setup HTTP routes
	http.HandleFunc("/transfer", handlers.TransferHandler(accountService))
    http.HandleFunc("/balance", handlers.BalanceHandler(accountService))
	
	// Start the server
	log.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
