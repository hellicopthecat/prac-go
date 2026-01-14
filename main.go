package main

import (
	"fmt"

	"github.com/hellicopthecat/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("choi")
	account.Deposit(5)
	fmt.Println(account.Balance())
	err := account.Withdraw(20)
	if err != nil {
		fmt.Println(err)
	}
	account.ChangeOwner("hohoho")
	fmt.Println(account.Balance(), account.GetOwner())
	fmt.Println(account)
}
