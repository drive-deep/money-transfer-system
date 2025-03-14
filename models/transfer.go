package models

// TransferRequest represents a money transfer request
type TransferRequest struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Amount int    `json:"amount"`
}
