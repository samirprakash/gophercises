package main

import (
	"errors"
)

// Wallet defines options for maintaining a wallet
type Wallet struct {
	balance Bitcoin
}

// Deposit adds bitcoins to wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Balance returns the total bitcoins in the wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// Withdraw removes bitcoin from wallet
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return errors.New("cannot withdraw, insufficient funds")
	}

	w.balance -= amount
	return nil
}
