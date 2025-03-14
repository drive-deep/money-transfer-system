package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	"github.com/stretchr/testify/assert"
)

// Test the balance retrieval functionality
func TestGetBalance(t *testing.T) {
	tests := []struct {
		user     string
		expected int
	}{
		{"Mark", 70},
		{"Jane", 50},
		{"Adam", 0},
	}

	for _, tt := range tests {
		t.Run("Get balance for "+tt.user, func(t *testing.T) {
			resp, err := http.Get("http://localhost:8080/balance?user=" + tt.user)
			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, resp.StatusCode)

			var result map[string]interface{}
			err = json.NewDecoder(resp.Body).Decode(&result)
			assert.NoError(t, err)

			assert.Equal(t, float64(tt.expected), result["balance"])
			assert.Equal(t, tt.user, result["user"])
		})
	}
}

// Test the money transfer functionality
func TestTransferMoney(t *testing.T) {
	tests := []struct {
		from     string
		to       string
		amount   int
		expected string
	}{
		{"Mark", "Jane", 30, "Transfer successful"},
		{"Jane", "Mark", 20, "Transfer successful"},
		{"Mark", "Mark", 10, "Error: Cannot transfer to your own account"},
		{"Mark", "Jane", 200, "Error: Insufficient funds"},
		{"InvalidUser", "Jane", 10, "Error: Invalid user"},
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
			assert.NoError(t, err)

			var result map[string]string
			err = json.NewDecoder(resp.Body).Decode(&result)
			assert.NoError(t, err)

			assert.Equal(t, tt.expected, result["message"])
		})
	}
}

// Test concurrency with multiple transfers happening at once
func TestConcurrency(t *testing.T) {
	// Simulate multiple transfers at once
	for i := 0; i < 10; i++ {
		go func(i int) {
			transferData := map[string]interface{}{
				"from":   "Mark",
				"to":     "Jane",
				"amount": 10,
			}
			data, err := json.Marshal(transferData)
			if err != nil {
				t.Errorf("Failed to marshal transfer data: %v", err)
				return
			}

			resp, err := http.Post("http://localhost:8080/transfer", "application/json", bytes.NewBuffer(data))
			if err != nil {
				t.Errorf("Failed to make transfer request: %v", err)
				return
			}
			defer resp.Body.Close()

			assert.Equal(t, http.StatusOK, resp.StatusCode)

			var result map[string]string
			err = json.NewDecoder(resp.Body).Decode(&result)
			if err != nil {
				t.Errorf("Failed to decode response: %v", err)
				return
			}

			assert.Equal(t, "Transfer successful", result["message"])
		}(i)
	}

	// Allow goroutines to complete
	// Typically use sync.WaitGroup to wait for goroutines in real tests
	// But for simplicity, we'll just wait here
	select {}
}
