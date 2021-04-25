// The package account simulates a bank account supporting:
// opening, closing, withdrawals, and deposits of money.
package account

import "sync"

type Account struct {
	closed  bool
	balance int64
	sync.RWMutex
}

// Open opens a bank account with 'initialDeposit' in it.
func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}

	return &Account{balance: initialDeposit}
}

// Close closes the bank account and return the balance.
func (a *Account) Close() (payout int64, ok bool) {
	a.Lock()
	defer a.Unlock()

	if a.closed {
		return 0, false
	}

	a.closed = true
	return a.balance, true
}

// Balance return the balance in the bank account.
func (a *Account) Balance() (balance int64, ok bool) {
	a.RLock()
	defer a.RUnlock()

	if a.closed {
		return 0, false
	}

	return a.balance, true
}

// Deposit makes deposits and withdrawals to the bank account
func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	a.Lock()
	defer a.Unlock()

	newBalance = a.balance + amount
	if a.closed || newBalance < 0 {
		return 0, false
	}

	a.balance = newBalance
	return newBalance, true
}
