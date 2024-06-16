package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin float64

// Make Bitcoin implement Stringer interface
func (b Bitcoin) String() string {
	return fmt.Sprintf("%g BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

// we need a pointer here to modify the real struct
func (w *Wallet) Deposit(amount Bitcoin) {
	// we could write (*w).balance, but it is not mandatory, in Go it is dereferenced implicitly
	w.balance += amount
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

// We don't need the pointer here, but it is a convention to keep consistency
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
