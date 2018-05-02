package main

import "testing"

func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, w Wallet, want Bitcoin) {
		t.Helper()
		got := w.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func(t *testing.T, got error, want error) {
		if got == nil {
			t.Fatal("wanted an error but did not get one!")
		}

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	assertNoError := func(t *testing.T, got error) {
		if got != nil {
			t.Fatal("got an error but did not want one")
		}
	}

	t.Run("add bitcoin to wallet", func(t *testing.T) {
		w := Wallet{}
		w.Deposit(Bitcoin(10))
		assertBalance(t, w, Bitcoin(10))
	})

	t.Run("withdraw bitcoin from wallet", func(t *testing.T) {
		w := Wallet{balance: Bitcoin(10)}
		err := w.Withdraw(Bitcoin(5))

		assertBalance(t, w, Bitcoin(5))
		assertNoError(t, err)
	})

	t.Run("withdraw insufficient funds from wallet", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		w := Wallet{balance: startingBalance}
		err := w.Withdraw(Bitcoin(100))

		assertBalance(t, w, startingBalance)
		assertError(t, err, ErrInsufficientFunds)
	})
}
