package test

import (
	"encoding/json"
	"net/http"
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestGetBalance(t *testing.T) {
	SetupTestDB()

	users := []string{"Mark", "Jane", "Adam"}


	// Now, perform assertions on the initial state
	for _, user := range users {
		t.Run("Get balance for "+user, func(t *testing.T) {
			resp, err := http.Get("http://localhost:8080/balance?user=" + user)
			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, resp.StatusCode)

			var result map[string]interface{}
			err = json.NewDecoder(resp.Body).Decode(&result)
			assert.NoError(t, err)
		})
	}
}
