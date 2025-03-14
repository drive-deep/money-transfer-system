package models

import "sync"

// Account represents a user account with balance and a mutex for concurrency safety
type Account struct {
	Balance int
	Mutex   sync.Mutex
}
