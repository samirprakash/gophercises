package main

import "testing"

func TestWallet(t *testing.T) {
	w := Wallet{}
	w.deposit(10)

	got := w.balance()
	want := 10
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
