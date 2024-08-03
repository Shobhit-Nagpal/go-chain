package block

import (
  "time"
  "github.com/Shobhit-Nagpal/go-chain/internal/txn"
)

type Metadata struct {
	Timestamp   time.Time
	BlockNumber int
	Hash        string
}

type Block struct {
	Metadata Metadata
	Data     []txn.Txn
}
