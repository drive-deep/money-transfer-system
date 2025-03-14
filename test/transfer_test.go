package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
)


// Test the money transfer functionality
func TestTransferMoney(t *testing.T) {
	SetupTestDB() // Ensure DB is initialized before tests run
	tests := []struct {
		from     string
		to       string
		amount   int
		expected string
	}{
		{"Mark", "Jane", 3, "Transfer successful"},
		{"Jane", "Mark", 20, "Transfer successful"},
		{"Mark", "Mark", 10, "Cannot transfer to the same account"},
		{"Mark", "Jane", 2000, "Insufficient funds"},
		{"InvalidUser", "Jane", 10, "Invalid user"},
	}

	for _, tt := range tests {
		t.Run("Transfer from "+tt.from+" to "+tt.to, func(t *testing.T) {
			transferData := map[string]interface{}{
				"from":   tt.from,
				"to":     tt.to,
				"amount": tt.amount,
			}
			data, err := json.Marshal(transferData)
			assert.NoError(t, err)

			resp, err := http.Post("http://localhost:8080/transfer", "application/json", bytes.NewBuffer(data))
			fmt.Println(resp, err)
			assert.NoError(t, err)

			var result map[string]string
			err = json.NewDecoder(resp.Body).Decode(&result)
			assert.NoError(t, err)
			fmt.Println("result")
            fmt.Println(result)
			assert.Equal(t, tt.expected, result["message"])
		})
	}
}