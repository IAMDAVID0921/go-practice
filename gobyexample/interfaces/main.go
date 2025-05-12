package main

import "fmt"

type IBankAccount interface {
	GetBalance() int
	Deposit(amount int)
	Withdraw(amount int) error
}

func main() {
	// wf := NewWellsFargo()
	// wf.Deposit(120)
	// wf.Deposit(240)
	// currentBalance := wf.GetBalance()
	// fmt.Printf("Current balance is: %d\n", currentBalance)
	// if err := wf.Withdraw(100); err != nil {
	// 	panic(err)
	// }
	// currentBalance = wf.GetBalance()
	// fmt.Printf("Current balance is: %d\n", currentBalance)
	// // if err := wf.Withdraw(300); err != nil {
	// // 	panic(err)
	// // }
	myAccounts := []IBankAccount{
		NewWellsFargo(),
		NewBitcoinAccount(),
	}

	for _, account := range myAccounts {
		account.Deposit(500)
		if err := account.Withdraw(70); err != nil {
			fmt.Printf("ERR: %d\n", err)
		}
		balance := account.GetBalance()
		fmt.Printf("balance is: %d\n", balance)
	}
}
