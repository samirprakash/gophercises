package main

import "testing"

func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, got, want Bitcoin) {
		t.Helper()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("add bitcoin to wallet", func(t *testing.T) {
		w := Wallet{}
		w.Deposit(Bitcoin(10))

		got := w.Balance()
		assertBalance(t, got, Bitcoin(10))
	})

	t.Run("withdraw bitcoin from wallet", func(t *testing.T) {
		w := Wallet{balance: Bitcoin(10)}
		w.Withdraw(Bitcoin(5))

		got := w.Balance()
		assertBalance(t, got, Bitcoin(5))
	})

}
