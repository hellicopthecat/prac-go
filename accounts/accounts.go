package accounts

import (
	"errors"
	"fmt"
)

// Account Struct
type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("Cant't Withdraw")

// New Account Create Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount f
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance Account
func (a Account) Balance() int {
	return a.balance
}

// Withdraw
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

// Change Owner
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

func (a Account) GetOwner() string {
	return a.owner
}

func (a Account) String() string {
	return fmt.Sprint(a.GetOwner(), "'s account. \nHas: ", a.Balance())
}
