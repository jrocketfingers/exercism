package account

import (
	"sync"
)

type Account struct {
	balance int64
	active  bool
    lock    sync.RWMutex
}

func Open(initialDeposit int64) *Account {
    if initialDeposit < 0 {
        return nil
    }
    acc := &Account{balance: initialDeposit, active: true}
	return acc
}

func (a *Account) Close() (payout int64, ok bool) {
    a.lock.Lock()
    defer a.lock.Unlock()
	ok = a.active
	payout = a.balance
    a.balance = 0
    a.active = false
	return
}

func (a *Account) Balance() (balance int64, ok bool) {
    a.lock.RLock()
    ok = a.active
    balance = a.balance
    a.lock.RUnlock()
	return
}

func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
    a.lock.Lock()
    defer a.lock.Unlock()

    if a.active != true {
        return a.balance, false
    }

    newBalance = a.balance + amount
    if newBalance < 0 {
        return newBalance, false
    } else {
        a.balance = newBalance
    }

    return a.balance, true
}
