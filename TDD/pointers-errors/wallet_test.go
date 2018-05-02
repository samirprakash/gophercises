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

	assertError := func(t *testing.T, got error, want string) {
		if got == nil {
			t.Fatal("wanted an error but did not get one!")
		}

		if got.Error() != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("add bitcoin to wallet", func(t *testing.T) {
		w := Wallet{}
		w.Deposit(Bitcoin(10))
		assertBalance(t, w, Bitcoin(10))
	})

	t.Run("withdraw bitcoin from wallet", func(t *testing.T) {
		w := Wallet{balance: Bitcoin(10)}
		w.Withdraw(Bitcoin(5))
		assertBalance(t, w, Bitcoin(5))
	})

	t.Run("withdraw insufficient funds from wallet", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		w := Wallet{balance: startingBalance}
		err := w.Withdraw(Bitcoin(100))

		assertBalance(t, w, startingBalance)
		assertError(t, err, "cannot withdraw, insufficient funds")
	})
}
