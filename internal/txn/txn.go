package txn

import (
	"github.com/Shobhit-Nagpal/go-chain/internal/account"
	"github.com/google/uuid"
)

const (
	SUCCESS = iota
	FAILURE
	PENDING
)

type Txn struct {
	Id       uuid.UUID
	Sender   *account.Account
	Receiver *account.Account
	Status   int
	Amount   float32
}

func CreateTransaction(sender, receiver *account.Account, amount float32) (*Txn, error) {
	id := uuid.New()
  transaction := &Txn{
		Id:       id,
		Sender:   sender,
		Receiver: receiver,
		Status:   PENDING,
    Amount: amount,
	}

  err := sender.Debit(amount)
  if err != nil {
    transaction.Status = FAILURE
    return transaction, err
  }

  receiver.Credit(amount)

  transaction.Status = SUCCESS

  return transaction, nil
}
