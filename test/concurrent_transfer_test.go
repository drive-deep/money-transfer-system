package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	"github.com/stretchr/testify/assert"
	"sync"
)

// Test concurrency with multiple transfers happening at once
func TestConcurrency(t *testing.T) {
	SetupTestDB() // Ensure DB is initialized before tests run
	var wg sync.WaitGroup
	numRequests := 10

	// Simulate multiple concurrent transfers
	for i := 0; i < numRequests; i++ {
		wg.Add(1) // Increment wait group counter

		go func(i int) {
			defer wg.Done() // Decrement counter when done

			transferData := map[string]interface{}{
				"from":   "Mark",
				"to":     "Jane",
				"amount": 1,
			}
			data, err := json.Marshal(transferData)
			assert.NoError(t, err, "Failed to marshal transfer data")

			resp, err := http.Post("http://localhost:8080/transfer", "application/json", bytes.NewBuffer(data))
			assert.NoError(t, err, "Failed to make transfer request")
			defer resp.Body.Close()

			assert.Equal(t, http.StatusOK, resp.StatusCode, "Unexpected status code")

			var result map[string]string
			err = json.NewDecoder(resp.Body).Decode(&result)
			assert.NoError(t, err, "Failed to decode response")

			assert.Equal(t, "Transfer successful", result["message"], "Unexpected response message")
		}(i)
	}

	wg.Wait() // Wait for all goroutines to finish
}