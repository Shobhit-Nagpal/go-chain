package block

import (
	"github.com/Shobhit-Nagpal/go-chain/internal/txn"
	"time"
)

type Metadata struct {
	Timestamp   time.Time
	BlockNumber int
	PrevHash    string
}

type Block struct {
	Metadata Metadata
	Data     []txn.Txn
}

func CreateBlock(prevHash string) Block {
	block := Block{
		Metadata: Metadata{
			Timestamp: time.Now(),
			PrevHash:  prevHash,
		},
		Data: []txn.Txn{},
	}

	return block
}
