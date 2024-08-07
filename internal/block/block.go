package block

import (
	"errors"
	"time"

	"github.com/Shobhit-Nagpal/go-chain/internal/err"
	"github.com/Shobhit-Nagpal/go-chain/internal/hash"
	"github.com/Shobhit-Nagpal/go-chain/internal/txn"
)

type Metadata struct {
	Timestamp   time.Time
	BlockNumber int
	PrevHash    []byte
	Hash        []byte
}

type Block struct {
	Metadata Metadata
	Data     map[string]*txn.Txn
}

func CreateBlock(prevHash []byte) *Block {
	blockHash := hash.CreateNewHash()

	block := &Block{
		Metadata: Metadata{
			Timestamp: time.Now(),
			PrevHash:  prevHash,
			Hash:      blockHash,
		},
		Data: map[string]*txn.Txn{},
	}

	return block
}

func (b *Block) GetCreatedAt() time.Time {
	return b.Metadata.Timestamp
}

func (b *Block) GetPreviousHash() string {
	return string(b.Metadata.PrevHash)
}

func (b *Block) GetHash() string {
	return string(b.Metadata.Hash)
}

func (b *Block) GetTxnInfo(txnId string) (txn.Txn, error) {
	transaction, exists := b.Data[txnId]
	if !exists {
		return *transaction, errors.New(err.TXN_NONEXISTENT)
	}

	return *transaction, nil
}

func (b *Block) AddTransaction(transaction *txn.Txn) error {
	_, exists := b.Data[transaction.Id.String()]
	if exists {
		return errors.New(err.TXN_EXISTS_IN_BLOCK)
	}

	b.Data[transaction.Id.String()] = transaction
	return nil
}
