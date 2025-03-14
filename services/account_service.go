package services

import (
	"fmt"
	"money-transfer-system/models"
	"sync"
)

// AccountService manages accounts and transfers
type AccountService struct {
	accounts map[string]*models.Account
	mutex    sync.Mutex
}

// NewAccountService initializes the account service with given accounts
func NewAccountService(accounts map[string]*models.Account) *AccountService {
	return &AccountService{
		accounts: accounts,
	}
}

// Transfer handles money transfer between accounts
func (s *AccountService) Transfer(req models.TransferRequest) error {
	if req.From == req.To {
		return fmt.Errorf("Cannot transfer to the same account")
	}
	if req.Amount <= 0 {
		return fmt.Errorf("Transfer amount must be greater than zero")
	}

	s.mutex.Lock()
	fromAcc, fromExists := s.accounts[req.From]
	toAcc, toExists := s.accounts[req.To]
	s.mutex.Unlock()

	if !fromExists || !toExists {
		return fmt.Errorf("Invalid user")
	}

	// Lock source account to prevent race conditions
	fromAcc.Mutex.Lock()
	defer fromAcc.Mutex.Unlock()

	// Check for sufficient balance
	if fromAcc.Balance < req.Amount {
		return fmt.Errorf("Insufficient funds")
	}

	// Lock destination account
	toAcc.Mutex.Lock()
	defer toAcc.Mutex.Unlock()

	// Perform the transfer
	fromAcc.Balance -= req.Amount
	toAcc.Balance += req.Amount

	return nil
}

func (s *AccountService) GetBalance(user string) (int, error) {
	s.mutex.Lock()
	account, exists := s.accounts[user]
	s.mutex.Unlock()

	if !exists {
		return 0, fmt.Errorf("user not found")
	}

	account.Mutex.Lock()
	balance := account.Balance
	account.Mutex.Unlock()

	return balance, nil
}