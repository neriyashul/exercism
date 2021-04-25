package account

import "sync"

type Account struct {
	open    bool
	balance int64
	mutex   sync.Mutex
}

func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}

	account := Account{
		open:    true,
		balance: initialDeposit,
	}

	return &account
}

func (a *Account) Close() (payout int64, ok bool) {
	if a != nil {
		a.mutex.Lock()
		defer a.mutex.Unlock()

		if a.isOpen() {
			a.open = false
			return a.balance, true
		}
	}

	return 0, false
}

func (a *Account) Balance() (balance int64, ok bool) {
	if a.isOpen() {
		return a.balance, true
	} else {
		return 0, false
	}
}

func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	if a != nil {
		a.mutex.Lock()
		defer a.mutex.Unlock()

		if a.isOpen() {
			newBalance := a.balance + amount
			if newBalance >= 0 {
				a.balance = newBalance
				return newBalance, true
			}
		}
	}

	return 0, false
}

func (a *Account) isOpen() bool {
	return a != nil && a.open
}
