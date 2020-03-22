package pointers

import (
	"errors"
	"fmt"
)

//Bitcoin type defined on top of int give us posibility to print out value with currencty BTC
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

//Wallet manage our budget
type Wallet struct {
	balance Bitcoin
}

//Deposit will add given amount to our current balance
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

//Balance will return current balance value
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

//ErrInsuficientFunds will be returned when there is not enough money on balance to withdraw
var ErrInsuficientFunds = errors.New("cannot witdraw, insufficient funds")

//Withdraw try to withdraw given value from our balance,
//returns an error when given value is bigger then balance.
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsuficientFunds
	}

	w.balance = w.balance - amount
	return nil
}
