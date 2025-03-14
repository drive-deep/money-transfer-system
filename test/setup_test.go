package test

import (
	"log"
	"money-transfer-system/models"
	"money-transfer-system/storage"
	"sync"
	"testing"
)

var testAccounts map[string]*models.Account
var once sync.Once

// SetupTestDB initializes the database once before running any tests.
func SetupTestDB() {
	once.Do(func() {
		log.Println("Initializing database for tests...")
		testAccounts = storage.InitializeDatabase()
	})
}

// TestMain is executed before any test runs in the `test` package
func TestMain(m *testing.M) {
	SetupTestDB() // Ensure database is initialized
	exitVal := m.Run()
	// No os.Exit() needed here since TestMain in multiple files may conflict
	// Instead, just return the exit value
	if exitVal != 0 {
		log.Fatalf("Tests failed with exit code: %d", exitVal)
	}
}
