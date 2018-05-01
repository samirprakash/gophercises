package main

// Wallet defines options for maintaining a wallet
type Wallet struct {
	balance int
}

// Deposit adds amount to the wallet
func (w *Wallet) Deposit(amount int) {
	w.balance += amount
}

// Balance returns the amount in a wallet
func (w *Wallet) Balance() int {
	return w.balance
}
