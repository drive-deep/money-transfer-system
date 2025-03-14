package storage

import (
	"money-transfer-system/models"
)

// InitializeDatabase sets up in-memory accounts with predefined balances
func InitializeDatabase() map[string]*models.Account {
	return map[string]*models.Account{
		"Mark": {Balance: 100},
		"Jane": {Balance: 50},
		"Adam": {Balance: 0},
	}
}

