package main

import "fmt"

// Stringer defines how we display Bitcoins
type Stringer interface {
	String() string
}

// Bitcoin is the type present in wallet
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet defines options for maintaining a wallet
type Wallet struct {
	balance Bitcoin
}

// Deposit adds amount to the wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Balance returns the amount in a wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
