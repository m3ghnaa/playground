package main

import "fmt"

//Bitcoin represents a number of Bitcoins
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

//Wallet stores the number of Bitcoin someone own
type Wallet struct {
	balance Bitcoin
}

//Deposit will add some Bitcoin to a wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

//Balance return the number of Bitcoins a wallet has
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
