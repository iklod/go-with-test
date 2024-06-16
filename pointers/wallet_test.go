package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
	// For more readability, helpers can be defined as func outside our function, so that we read directly the tests and not some helpers
	checkBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}
	checkError := func(t testing.TB, got, want error) {
		t.Helper()
		if got == nil {
			// Fatal means the code execution will stop here. Otherwise we would use a nil pointer
			t.Fatal("we wanted an error here")
		}
		if got != want {
			t.Errorf("want %q got %q", want, got)
		}
	}
	checkNoError := func(t testing.TB, got error) {
		t.Helper()
		if got != nil {
			t.Fatal("got an error but didn't want one")
		}
	}
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(10)

		want := Bitcoin(10)
		checkBalance(t, wallet, want)
	})
	t.Run("withdraw with funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))

		checkNoError(t, err)
		checkBalance(t, wallet, Bitcoin(10))
	})
	t.Run("Withdraw unsufficient found", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(40)
		checkError(t, err, ErrInsufficientFunds)
		checkBalance(t, wallet, startingBalance)
	})
}
