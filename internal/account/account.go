package account

import (
	"errors"

	"github.com/Shobhit-Nagpal/go-chain/internal/err"
	"github.com/google/uuid"
)

type Account struct {
	Id      uuid.UUID
	Balance float32
}

func CreateAccount() *Account {
	return &Account{
		Id:      uuid.New(),
		Balance: 100.00,
	}
}

func (ac *Account) GetBalance() float32 {
  return ac.Balance
}

func (ac *Account) Debit(amount float32) error {
  if ac.Balance < amount {
    return errors.New(err.NOT_ENOUGH_BALANCE)
  }

  ac.Balance = ac.Balance - amount
  return nil
}

func (ac *Account) Credit(amount float32) {
  ac.Balance = ac.Balance + amount
}
