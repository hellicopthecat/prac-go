package main

import (
	"fmt"

	"github.com/hellicopthecat/learngo/dict"
)

func main() {
	// account := accounts.NewAccount("choi")
	// account.Deposit(5)
	// fmt.Println(account.Balance())
	// err := account.Withdraw(20)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// account.ChangeOwner("hohoho")
	// fmt.Println(account.Balance(), account.GetOwner())
	// fmt.Println(account)

	myDict := dict.Dictionary{"first": "First word"}
	// myDict["hoho"] = "haha"
	// fmt.Println(myDict)
	definition, err := myDict.Search("first")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
	myDict.Add("hello", "definitions")
	definition2, _ := myDict.Search("hello")
	fmt.Println(definition2)
	myDict.Update("hello", "This is say hello")
	definition3, _ := myDict.Search("hello")
	fmt.Println(definition3)
}
