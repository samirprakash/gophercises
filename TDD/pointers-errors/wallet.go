package main

// Wallet defines options for maintaining a wallet
type Wallet struct {
	Amount int
}

func (w *Wallet) deposit(a int) {
	w.Amount = a
}

func (w *Wallet) balance() int {
	return w.Amount
}
