package txn

import (
	"errors"
	"time"

	"github.com/Shobhit-Nagpal/go-chain/internal/account"
	"github.com/Shobhit-Nagpal/go-chain/internal/err"
	"github.com/google/uuid"
)

const (
	SUCCESS = iota
	FAILURE
	PENDING
)

type Txn struct {
	Id        uuid.UUID
	Sender    *account.Account
	Receiver  *account.Account
	Status    int
	Amount    float32
	Timestamp time.Time
}

func CreateTransactionPool() []*Txn {
	return []*Txn{}
}

func EnqueueTxn(pool []*Txn, txn *Txn) []*Txn {
  pool = append(pool, txn)
  return pool
}

func DequeueTxn(pool []*Txn) ([]*Txn, *Txn, error) {

  if len(pool) == 0 {
    return nil, nil, errors.New(err.EMPTY_QUEUE_DEQUEUE_ATTEMPT)
  }

  dequeued := pool[0]
  if len(pool) == 1 {
    return []*Txn{}, dequeued, nil
  }

  return pool[1:], dequeued, nil
}

func CreateTransaction(sender, receiver *account.Account, amount float32) (*Txn, error) {
	id := uuid.New()
	transaction := &Txn{
		Id:        id,
		Sender:    sender,
		Receiver:  receiver,
		Status:    PENDING,
		Amount:    amount,
		Timestamp: time.Now(),
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

func (t *Txn) CreatedAt() time.Time {
	return t.Timestamp
}
