package main

import "errors"

type BitcoinAccount struct {
	balance int
	fee     int
}

func NewBitcoinAccount() *BitcoinAccount {
	return &BitcoinAccount{
		balance: 0,
		fee:     300,
	}
}

func (b *BitcoinAccount) GetBalance() int {
	return b.balance
}

func (b *BitcoinAccount) Deposit(amount int) {
	b.balance += amount
}

func (b *BitcoinAccount) Withdraw(amount int) error {
	NewBalance := b.balance - amount - b.fee
	if NewBalance < 0 {
		return errors.New("insufficient funds")
	}
	b.balance = NewBalance
	return nil
}
