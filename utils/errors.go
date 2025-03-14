package utils

import "errors"

var (
	ErrInvalidAccount = errors.New("invalid account")
	ErrInsufficientFunds = errors.New("insufficient funds")
	ErrSelfTransfer = errors.New("cannot transfer to the same account")
	ErrInvalidAmount = errors.New("transfer amount must be greater than zero")
)
